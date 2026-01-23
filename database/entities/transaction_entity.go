package entities

import (
	"time"
)

type PaymentStatus string

const (
	PaymentStatusPending PaymentStatus = "pending"
	PaymentStatusPaid    PaymentStatus = "paid"
	PaymentStatusExpired PaymentStatus = "expired"
)

type TransactionStatus string

const (
	TransactionStatusConfirmed TransactionStatus = "confirmed"
	TransactionStatusCompleted TransactionStatus = "completed"
	TransactionStatusCancelled TransactionStatus = "cancelled"
)

const PaymentDeadlineMinutes = 15

type Transaction struct {
	ID                uint              `gorm:"primaryKey;autoIncrement" json:"id"`
	MitraID           uint              `gorm:"not null;index" json:"mitra_id"`
	ReceiptID         uint              `gorm:"not null;index" json:"receipt_id"`
	TransactionNumber string            `gorm:"type:varchar(50);uniqueIndex;not null" json:"transaction_number"`
	Subtotal          int               `gorm:"not null" json:"subtotal"`
	TotalAmount       int               `gorm:"not null" json:"total_amount"`
	PaymentAmount     int               `gorm:"not null" json:"payment_amount"`
	PaymentStatus     PaymentStatus     `gorm:"type:varchar(20);not null;default:'pending'" json:"payment_status"`
	PaymentDeadline   time.Time         `gorm:"type:timestamp with time zone;not null" json:"payment_deadline"`
	PaymentReference  string            `gorm:"type:varchar(255)" json:"payment_reference"`
	Status            TransactionStatus `gorm:"type:varchar(20);not null;default:'confirmed'" json:"status"`
	Notes             string            `gorm:"type:text" json:"notes"`
	PickupDate        *time.Time        `gorm:"type:date" json:"pickup_date"`
	QRCode            string            `gorm:"type:varchar(100);uniqueIndex" json:"qr_code"`
	QRCodeScanned     bool              `gorm:"default:false" json:"qr_code_scanned"`
	QRScannedAt       *time.Time        `gorm:"type:timestamp with time zone" json:"qr_scanned_at"`
	PickupTimeStart   string            `gorm:"type:time" json:"pickup_time_start"`
	PickupTimeEnd     string            `gorm:"type:time" json:"pickup_time_end"`
	CompletedAt       *time.Time        `gorm:"type:timestamp with time zone" json:"completed_at"`

	Mitra            MitraProfile      `gorm:"foreignKey:MitraID" json:"mitra,omitempty"`
	Receipt          ReceiptProfile    `gorm:"foreignKey:ReceiptID" json:"receipt,omitempty"`
	TransactionItems []TransactionItem `gorm:"foreignKey:TransactionID" json:"transaction_items,omitempty"`

	Timestamp
}