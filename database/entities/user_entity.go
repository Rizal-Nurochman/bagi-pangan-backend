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
	ID               uint       `gorm:"primaryKey;autoIncrement"`
	Email            string     `gorm:"type:varchar(255);uniqueIndex;not null"`
	Phone            string     `gorm:"type:varchar(20);uniqueIndex;not null"`
	PasswordHash     string     `gorm:"type:varchar(255);not null"`
	Role             UserRole   `gorm:"type:varchar(50);not null;default:'penerima'"`
	CodeVerification string     `gorm:"type:varchar(10)"`
	CodeExpiredAt    *time.Time `gorm:"type:timestamp with time zone"`
	EmailVerified    bool       `gorm:"default:false"`

	MitraProfile   *MitraProfile   `gorm:"foreignKey:UserID"`
	ReceiptProfile *ReceiptProfile `gorm:"foreignKey:UserID"`

	Timestamp
}
