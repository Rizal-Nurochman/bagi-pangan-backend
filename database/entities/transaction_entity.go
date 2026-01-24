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
	ID                uint              `gorm:"primaryKey;autoIncrement"`
	MitraID           uint              `gorm:"not null;index"`
	ReceiptID         uint              `gorm:"not null;index"`
	TransactionNumber string            `gorm:"type:varchar(50);uniqueIndex;not null"`
	Subtotal          int               `gorm:"not null"`
	TotalAmount       int               `gorm:"not null"`
	PaymentAmount     int               `gorm:"not null"`
	PaymentStatus     PaymentStatus     `gorm:"type:varchar(20);not null;default:'pending'"`
	PaymentDeadline   time.Time         `gorm:"type:timestamp with time zone;not null"`
	PaymentReference  string            `gorm:"type:varchar(255)"`
	Status            TransactionStatus `gorm:"type:varchar(20);not null;default:'confirmed'"`
	Notes             string            `gorm:"type:text"`
	PickupDate        *time.Time        `gorm:"type:date"`
	QRCode            string            `gorm:"type:varchar(100);uniqueIndex"`
	QRCodeScanned     bool              `gorm:"default:false"`
	QRScannedAt       *time.Time        `gorm:"type:timestamp with time zone"`
	PickupTimeStart   string            `gorm:"type:time"`
	PickupTimeEnd     string            `gorm:"type:time"`
	CompletedAt       *time.Time        `gorm:"type:timestamp with time zone"`

	Mitra            MitraProfile      `gorm:"foreignKey:MitraID"`
	Receipt          ReceiptProfile    `gorm:"foreignKey:ReceiptID"`
	TransactionItems []TransactionItem `gorm:"foreignKey:TransactionID"`
	Timestamp
}