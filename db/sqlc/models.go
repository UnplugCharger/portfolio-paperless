// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package db

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"
)

type AuditInvoice struct {
	AuditID             int64      `json:"audit_id"`
	Operation           *string    `json:"operation"`
	ID                  *int32     `json:"id"`
	InvoiceNo           *string    `json:"invoice_no"`
	PurchaseOrderNumber *string    `json:"purchase_order_number"`
	Date                *time.Time `json:"date"`
	Attn                *string    `json:"attn"`
	CompanyID           *int32     `json:"company_id"`
	Site                *string    `json:"site"`
	AmountDue           *float64   `json:"amount_due"`
	BankDetails         *int32     `json:"bank_details"`
	SignatoryID         *int32     `json:"signatory_id"`
	SentOrReceived      *string    `json:"sent_or_received"`
	ChangedAt           time.Time  `json:"changed_at"`
	ChangedBy           *int32     `json:"changed_by"`
}

type AuditPaymentRequest struct {
	AuditID      int64      `json:"audit_id"`
	Operation    *string    `json:"operation"`
	RequestID    *int32     `json:"request_id"`
	EmployeeID   *int32     `json:"employee_id"`
	Amount       *float64   `json:"amount"`
	Description  *string    `json:"description"`
	RequestDate  *time.Time `json:"request_date"`
	Status       *string    `json:"status"`
	AdminID      *int32     `json:"admin_id"`
	ApprovalDate *time.Time `json:"approval_date"`
	InvoiceID    *int32     `json:"invoice_id"`
	ChangedAt    time.Time  `json:"changed_at"`
	ChangedBy    *int32     `json:"changed_by"`
}

type AuditPettyCash struct {
	AuditID         int64      `json:"audit_id"`
	Operation       *string    `json:"operation"`
	TransactionID   *int32     `json:"transaction_id"`
	EmployeeID      *int32     `json:"employee_id"`
	Amount          *string    `json:"amount"`
	Description     *string    `json:"description"`
	TransactionDate *time.Time `json:"transaction_date"`
	UpdatedAt       *time.Time `json:"updated_at"`
	Folio           *string    `json:"folio"`
	DebitAccount    *string    `json:"debit_account"`
	Status          *string    `json:"status"`
	AuthorisedBy    *int32     `json:"authorised_by"`
	ChangedAt       time.Time  `json:"changed_at"`
	ChangedBy       *int32     `json:"changed_by"`
}

type AuditPurchaseOrder struct {
	AuditID        int64      `json:"audit_id"`
	Operation      *string    `json:"operation"`
	ID             *int32     `json:"id"`
	PoNo           *string    `json:"po_no"`
	Date           *time.Time `json:"date"`
	Attn           *string    `json:"attn"`
	CompanyID      *int32     `json:"company_id"`
	Address        *string    `json:"address"`
	SignatoryID    *int32     `json:"signatory_id"`
	QuotationID    *int32     `json:"quotation_id"`
	PoStatus       *string    `json:"po_status"`
	SentOrReceived *string    `json:"sent_or_received"`
	ChangedAt      time.Time  `json:"changed_at"`
	ChangedBy      *int32     `json:"changed_by"`
}

type AuditQuotation struct {
	AuditID                 int64      `json:"audit_id"`
	Operation               *string    `json:"operation"`
	ID                      *int32     `json:"id"`
	QuotationNo             *string    `json:"quotation_no"`
	QuotationRevision       *int32     `json:"quotation_revision"`
	QuotationRevisionNumber *int32     `json:"quotation_revision_number"`
	Date                    *time.Time `json:"date"`
	Attn                    *string    `json:"attn"`
	CompanyID               *int32     `json:"company_id"`
	Site                    *string    `json:"site"`
	Validity                *int32     `json:"validity"`
	Warranty                *int32     `json:"warranty"`
	PaymentTerms            *string    `json:"payment_terms"`
	DeliveryTerms           *string    `json:"delivery_terms"`
	SignatoryID             *int32     `json:"signatory_id"`
	Status                  *string    `json:"status"`
	SentOrReceived          *string    `json:"sent_or_received"`
	ChangedAt               time.Time  `json:"changed_at"`
	ChangedBy               *int32     `json:"changed_by"`
}

type BankDetail struct {
	ID            int32  `json:"id"`
	BankName      string `json:"bank_name"`
	AccountName   string `json:"account_name"`
	AccountNumber string `json:"account_number"`
	Branch        string `json:"branch"`
	SwiftCode     string `json:"swift_code"`
	Address       string `json:"address"`
	Country       string `json:"country"`
	Currency      string `json:"currency"`
	AccountType   string `json:"account_type"`
	CompanyID     int32  `json:"company_id"`
}

type Car struct {
	ID           int32     `json:"id"`
	Make         string    `json:"make"`
	Model        string    `json:"model"`
	Year         int32     `json:"year"`
	LicensePlate string    `json:"license_plate"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type CarFuelConsumption struct {
	ConsumptionID int32     `json:"consumption_id"`
	CarID         int32     `json:"car_id"`
	LitersOfFuel  *float64  `json:"liters_of_fuel"`
	CostInKsh     float64   `json:"cost_in_ksh"`
	FuelDate      time.Time `json:"fuel_date"`
	TransactionID *int32    `json:"transaction_id"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Company struct {
	ID       int64   `json:"id"`
	Name     string  `json:"name"`
	Initials string  `json:"initials"`
	Address  *string `json:"address"`
}

type Invoice struct {
	ID                  int32      `json:"id"`
	InvoiceNo           string     `json:"invoice_no"`
	PurchaseOrderNumber string     `json:"purchase_order_number"`
	Date                *time.Time `json:"date"`
	Attn                string     `json:"attn"`
	CompanyID           int32      `json:"company_id"`
	Site                string     `json:"site"`
	AmountDue           float64    `json:"amount_due"`
	BankDetails         int32      `json:"bank_details"`
	SignatoryID         int32      `json:"signatory_id"`
	SentOrReceived      string     `json:"sent_or_received"`
	PdfUrl              *string    `json:"pdf_url"`
}

type InvoiceItem struct {
	ID          int32   `json:"id"`
	Description string  `json:"description"`
	Uom         string  `json:"uom"`
	Qty         int32   `json:"qty"`
	UnitPrice   float64 `json:"unit_price"`
	NetPrice    float64 `json:"net_price"`
	Currency    string  `json:"currency"`
	InvoiceID   int32   `json:"invoice_id"`
}

type PaymentRequest struct {
	RequestID        int32      `json:"request_id"`
	PaymentRequestNo string     `json:"payment_request_no"`
	AmountInWords    string     `json:"amount_in_words"`
	EmployeeID       int32      `json:"employee_id"`
	Currency         string     `json:"currency"`
	Amount           float64    `json:"amount"`
	Description      string     `json:"description"`
	RequestDate      time.Time  `json:"request_date"`
	Status           string     `json:"status"`
	AdminID          *int32     `json:"admin_id"`
	ApprovalDate     *time.Time `json:"approval_date"`
	InvoiceID        *int32     `json:"invoice_id"`
	PdfUrl           *string    `json:"pdf_url"`
}

type PettyCash struct {
	TransactionID   int32          `json:"transaction_id"`
	PettyCashNo     string         `json:"petty_cash_no"`
	EmployeeID      int32          `json:"employee_id"`
	Amount          pgtype.Numeric `json:"amount"`
	CurrencyCode    *string        `json:"currency_code"`
	Description     string         `json:"description"`
	TransactionDate time.Time      `json:"transaction_date"`
	UpdatedAt       time.Time      `json:"updated_at"`
	ApprovedAt      *time.Time     `json:"approved_at"`
	Folio           string         `json:"folio"`
	DebitAccount    string         `json:"debit_account"`
	Status          string         `json:"status"`
	AuthorisedBy    *int32         `json:"authorised_by"`
	PdfUrl          *string        `json:"pdf_url"`
}

type PurchaseOrder struct {
	ID             int32      `json:"id"`
	PoNo           string     `json:"po_no"`
	Date           *time.Time `json:"date"`
	Attn           string     `json:"attn"`
	CompanyID      int32      `json:"company_id"`
	Address        string     `json:"address"`
	SignatoryID    int32      `json:"signatory_id"`
	QuotationID    *int32     `json:"quotation_id"`
	PoStatus       *string    `json:"po_status"`
	ApprovedBy     *int32     `json:"approved_by"`
	SentOrReceived *string    `json:"sent_or_received"`
	PdfUrl         *string    `json:"pdf_url"`
}

type PurchaseOrderItem struct {
	ID              int32   `json:"id"`
	Description     string  `json:"description"`
	PartNo          string  `json:"part_no"`
	Uom             string  `json:"uom"`
	Qty             int32   `json:"qty"`
	ItemPrice       float64 `json:"item_price"`
	Discount        float64 `json:"discount"`
	NetPrice        float64 `json:"net_price"`
	NetValue        float64 `json:"net_value"`
	Currency        string  `json:"currency"`
	PurchaseOrderID int32   `json:"purchase_order_id"`
}

type Quotation struct {
	ID                      int32      `json:"id"`
	QuotationNo             string     `json:"quotation_no"`
	QuotationRevision       int32      `json:"quotation_revision"`
	QuotationRevisionNumber int32      `json:"quotation_revision_number"`
	Date                    *time.Time `json:"date"`
	Attn                    string     `json:"attn"`
	CompanyID               int32      `json:"company_id"`
	Site                    string     `json:"site"`
	Validity                int32      `json:"validity"`
	Warranty                int32      `json:"warranty"`
	PaymentTerms            string     `json:"payment_terms"`
	DeliveryTerms           string     `json:"delivery_terms"`
	SignatoryID             int32      `json:"signatory_id"`
	Status                  string     `json:"status"`
	SentOrReceived          *string    `json:"sent_or_received"`
	PdfUrl                  *string    `json:"pdf_url"`
}

type QuotationItem struct {
	ID          int32   `json:"id"`
	Description string  `json:"description"`
	Uom         string  `json:"uom"`
	Qty         int32   `json:"qty"`
	LeadTime    string  `json:"lead_time"`
	ItemPrice   float64 `json:"item_price"`
	Disc        float64 `json:"disc"`
	UnitPrice   float64 `json:"unit_price"`
	NetPrice    float64 `json:"net_price"`
	Currency    string  `json:"currency"`
	QuotationID int32   `json:"quotation_id"`
}

type Role struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type Session struct {
	ID           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiresAt    time.Time `json:"expires_at"`
	CreatedAt    time.Time `json:"created_at"`
}

type Signatory struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Title string `json:"title"`
}

type User struct {
	ID                int64     `json:"id"`
	Username          string    `json:"username"`
	HashedPassword    string    `json:"hashed_password"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	Department        string    `json:"department"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

type UserRole struct {
	ID           int64      `json:"id"`
	UserID       int64      `json:"user_id"`
	RoleID       int64      `json:"role_id"`
	CreatedAt    time.Time  `json:"created_at"`
	TerminatedAt *time.Time `json:"terminated_at"`
}
