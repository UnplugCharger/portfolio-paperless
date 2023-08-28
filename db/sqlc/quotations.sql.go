// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: quotations.sql

package db

import (
	"context"
)

const createQuotation = `-- name: CreateQuotation :one
INSERT INTO quotations ( attn, company_id, site, validity, warranty, payment_terms, delivery_terms, signatory_id, status, sent_or_received)
VALUES ( $1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING id, quotation_no, quotation_revision, quotation_revision_number, date, attn, company_id, site, validity, warranty, payment_terms, delivery_terms, signatory_id, status, sent_or_received, pdf_url
`

type CreateQuotationParams struct {
	Attn           string  `json:"attn"`
	CompanyID      int32   `json:"company_id"`
	Site           string  `json:"site"`
	Validity       int32   `json:"validity"`
	Warranty       int32   `json:"warranty"`
	PaymentTerms   string  `json:"payment_terms"`
	DeliveryTerms  string  `json:"delivery_terms"`
	SignatoryID    int32   `json:"signatory_id"`
	Status         string  `json:"status"`
	SentOrReceived *string `json:"sent_or_received"`
}

func (q *Queries) CreateQuotation(ctx context.Context, arg CreateQuotationParams) (Quotation, error) {
	row := q.db.QueryRow(ctx, createQuotation,
		arg.Attn,
		arg.CompanyID,
		arg.Site,
		arg.Validity,
		arg.Warranty,
		arg.PaymentTerms,
		arg.DeliveryTerms,
		arg.SignatoryID,
		arg.Status,
		arg.SentOrReceived,
	)
	var i Quotation
	err := row.Scan(
		&i.ID,
		&i.QuotationNo,
		&i.QuotationRevision,
		&i.QuotationRevisionNumber,
		&i.Date,
		&i.Attn,
		&i.CompanyID,
		&i.Site,
		&i.Validity,
		&i.Warranty,
		&i.PaymentTerms,
		&i.DeliveryTerms,
		&i.SignatoryID,
		&i.Status,
		&i.SentOrReceived,
		&i.PdfUrl,
	)
	return i, err
}

const getQuotationByID = `-- name: GetQuotationByID :one
SELECT id, quotation_no, quotation_revision, quotation_revision_number, date, attn, company_id, site, validity, warranty, payment_terms, delivery_terms, signatory_id, status, sent_or_received, pdf_url FROM quotations
WHERE id = $1
`

func (q *Queries) GetQuotationByID(ctx context.Context, id int32) (Quotation, error) {
	row := q.db.QueryRow(ctx, getQuotationByID, id)
	var i Quotation
	err := row.Scan(
		&i.ID,
		&i.QuotationNo,
		&i.QuotationRevision,
		&i.QuotationRevisionNumber,
		&i.Date,
		&i.Attn,
		&i.CompanyID,
		&i.Site,
		&i.Validity,
		&i.Warranty,
		&i.PaymentTerms,
		&i.DeliveryTerms,
		&i.SignatoryID,
		&i.Status,
		&i.SentOrReceived,
		&i.PdfUrl,
	)
	return i, err
}

const listQuotations = `-- name: ListQuotations :many
SELECT id, quotation_no, quotation_revision, quotation_revision_number, date, attn, company_id, site, validity, warranty, payment_terms, delivery_terms, signatory_id, status, sent_or_received, pdf_url FROM quotations
LIMIT $1 OFFSET $2
`

type ListQuotationsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListQuotations(ctx context.Context, arg ListQuotationsParams) ([]Quotation, error) {
	rows, err := q.db.Query(ctx, listQuotations, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Quotation{}
	for rows.Next() {
		var i Quotation
		if err := rows.Scan(
			&i.ID,
			&i.QuotationNo,
			&i.QuotationRevision,
			&i.QuotationRevisionNumber,
			&i.Date,
			&i.Attn,
			&i.CompanyID,
			&i.Site,
			&i.Validity,
			&i.Warranty,
			&i.PaymentTerms,
			&i.DeliveryTerms,
			&i.SignatoryID,
			&i.Status,
			&i.SentOrReceived,
			&i.PdfUrl,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateQuotation = `-- name: UpdateQuotation :one
UPDATE quotations
SET attn = COALESCE($1, attn),
    company_id = COALESCE($2, company_id),
    site = COALESCE($3, site),
    validity = COALESCE($4, validity),
    warranty = COALESCE($5, warranty),
    payment_terms = COALESCE($6, payment_terms),
    delivery_terms = COALESCE($7, delivery_terms),
    signatory_id = COALESCE($8, signatory_id),
    status = COALESCE($9, status),
    sent_or_received = COALESCE($10, sent_or_received)
WHERE id = $11
RETURNING id, quotation_no, quotation_revision, quotation_revision_number, date, attn, company_id, site, validity, warranty, payment_terms, delivery_terms, signatory_id, status, sent_or_received, pdf_url
`

type UpdateQuotationParams struct {
	Attn           string  `json:"attn"`
	CompanyID      int32   `json:"company_id"`
	Site           string  `json:"site"`
	Validity       int32   `json:"validity"`
	Warranty       int32   `json:"warranty"`
	PaymentTerms   string  `json:"payment_terms"`
	DeliveryTerms  string  `json:"delivery_terms"`
	SignatoryID    int32   `json:"signatory_id"`
	Status         string  `json:"status"`
	SentOrReceived *string `json:"sent_or_received"`
	ID             int32   `json:"id"`
}

func (q *Queries) UpdateQuotation(ctx context.Context, arg UpdateQuotationParams) (Quotation, error) {
	row := q.db.QueryRow(ctx, updateQuotation,
		arg.Attn,
		arg.CompanyID,
		arg.Site,
		arg.Validity,
		arg.Warranty,
		arg.PaymentTerms,
		arg.DeliveryTerms,
		arg.SignatoryID,
		arg.Status,
		arg.SentOrReceived,
		arg.ID,
	)
	var i Quotation
	err := row.Scan(
		&i.ID,
		&i.QuotationNo,
		&i.QuotationRevision,
		&i.QuotationRevisionNumber,
		&i.Date,
		&i.Attn,
		&i.CompanyID,
		&i.Site,
		&i.Validity,
		&i.Warranty,
		&i.PaymentTerms,
		&i.DeliveryTerms,
		&i.SignatoryID,
		&i.Status,
		&i.SentOrReceived,
		&i.PdfUrl,
	)
	return i, err
}
