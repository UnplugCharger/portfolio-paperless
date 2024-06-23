package main

import (
	"context"
	"database/sql"
	"errors"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/hibiken/asynq"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/qwetu_petro/backend/workers"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"github.com/qwetu_petro/backend/gapi"
	"github.com/qwetu_petro/backend/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	db "github.com/qwetu_petro/backend/db/sqlc"
	"github.com/qwetu_petro/backend/utils"

	"github.com/qwetu_petro/backend/api"

	"github.com/qwetu_petro/backend/logger"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v4/stdlib"
)

var log = logger.NewLogger()

func main() {

	log := logger.NewLogger()
	conf, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("could not load config")
	}

	// Use pgxpool for connection pooling
	poolConfig, err := pgxpool.ParseConfig(conf.DBSource)
	if err != nil {
		log.Fatal("Error parsing pool config")
	}

	// You can set pool configuration options here if needed
	poolConfig.MaxConns = 20 // For example, setting the max number of connections in the pool

	connPool, err := pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		log.Fatal("Error connecting to the database")
	}

	log.Debug("migration string " + conf.MigrationUrl)

	// Assuming your runDbMigrations accepts a pgxpool
	runDbMigrations(conf.MigrationUrl, conf.DBSource)

	// Pass the connection pool instead of a single connection
	store := db.NewStore(connPool)

	redisOpts := asynq.RedisClientOpt{
		Addr:     conf.RedisAddress,
		Password: conf.RedisPassword,
		//Username: conf.RedisUsername,
	}

	taskDistributor := workers.NewRedisTaskDistributor(redisOpts)

	// Assuming your runTaskProcessor and runGinServer accept a store with pgxpool
	go runTaskProcessor(store, redisOpts)
	go runGrpcServer(conf, store)
	go runGateWayServer(conf, store)
	runGinServer(conf, store, taskDistributor)
}
func runGinServer(config utils.Config, store db.Store, taskDistributor workers.TaskDistributor) {
	server, err := api.NewServer(config, store, taskDistributor)
	if err != nil {
		log.Fatal("can not create server")
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("can not start server")
	}
}

func runGrpcServer(config utils.Config, store db.Store) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("can not create server")
	}

	grpcServer := grpc.NewServer()
	pb.RegisterQwetuBackendGrpcServiceServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("can not start server")
	}
	log.Info("Starting gRPC server at " + config.GRPCServerAddress)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("can not start  GRPC server")
	}

}

// Grpc Gateway Server set Up
func runGateWayServer(config utils.Config, store db.Store) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("can not create server")
	}
	grpcMux := runtime.NewServeMux()
	cxt, cancel := context.WithCancel(context.Background())
	defer cancel()
	err = pb.RegisterQwetuBackendGrpcServiceHandlerServer(cxt, grpcMux, server)
	if err != nil {
		log.Fatal("can not create server")
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	listener, err := net.Listen("tcp", config.HTTPServerAddress)
	if err != nil {
		log.Fatal("can not start server")
	}
	log.Info("Starting HTTP gateway server at " + config.GRPCServerAddress)
	err = http.Serve(listener, mux)
	if err != nil {
		log.Fatal("can not start  HTTP server")
	}

}

func runDbMigrations(path string, dbSource string) {
	// Use pgx as a driver for database/sql
	dB, err := sql.Open("pgx", dbSource)
	if err != nil {
		log.Fatal("failed to create *sql.DB instance")
	}
	defer dB.Close()

	driver, err := postgres.WithInstance(dB, &postgres.Config{})
	if err != nil {
		log.Fatal("failed to create migration driver instance")
	}

	migration, err := migrate.NewWithDatabaseInstance(path, "postgres", driver)
	if err != nil {
		log.Fatal("failed to create migration instance")
	}

	if err = migration.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal("failed to run migration")
	}

	log.Info("Database migration successful")
}

func runTaskProcessor(store db.Store, redisOpts asynq.RedisClientOpt) {
	taskProcessor := workers.NewRedisTaskProcessor(redisOpts, store)
	log.Info("Starting task processor ")

	err := taskProcessor.Start()
	if err != nil {
		log.Fatal("failed  to start task processor")
	}

}
