package gapi

import (
	"context"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

const (
	grpcGateWayUserAgentHeader = "grpcgateway-user-agent"
	xForwardedForHeader        = "x-forwarded-for"
	userAgentHeader            = "userAgent"
)

type Metadata struct {
	UserAgent string
	ClientIP  string
}

// extract grpc metadata from context
func (server *Server) extractMetadata(ctx context.Context) Metadata {
	mtdt := Metadata{}
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if userAgents := md.Get(grpcGateWayUserAgentHeader); len(userAgents) > 0 {
			mtdt.UserAgent = userAgents[0]
		}
		if userAgents := md.Get(userAgentHeader); len(userAgents) > 0 {
			mtdt.UserAgent = userAgents[0]
		}
		if clientIps := md.Get(xForwardedForHeader); len(clientIps) > 0 {
			mtdt.ClientIP = clientIps[0]
		}
	}
	if p, ok := peer.FromContext(ctx); ok {
		mtdt.ClientIP = p.Addr.String()
	}

	return mtdt
}
