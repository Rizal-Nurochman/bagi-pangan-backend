package entities

import (
	"time"

)

type UserRole string

const (
	UserRoleMitra    UserRole = "mitra"
	UserRolePenerima UserRole = "penerima"
	UserRoleAdmin    UserRole = "admin"
)

type User struct {
	ID               uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Email            string     `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	Phone            string     `gorm:"type:varchar(20);uniqueIndex;not null" json:"phone"`
	PasswordHash     string     `gorm:"type:varchar(255);not null" json:"-"`
	Role             UserRole   `gorm:"type:varchar(50);not null;default:'penerima'" json:"role"`
	CodeVerification string     `gorm:"type:varchar(10)" json:"-"`
	CodeExpiredAt    *time.Time `gorm:"type:timestamp with time zone"`
	EmailVerified    bool       `gorm:"default:false" json:"email_verified"`
	LastLoginAt      *time.Time `gorm:"type:timestamp with time zone" json:"last_login_at"`

	MitraProfile   *MitraProfile   `gorm:"foreignKey:UserID" json:"mitra_profile,omitempty"`
	ReceiptProfile *ReceiptProfile `gorm:"foreignKey:UserID" json:"receipt_profile,omitempty"`

	Timestamp
}
