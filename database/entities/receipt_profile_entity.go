package entities

import (
	"time"
)

type ReceiptProfile struct {
	ID          uint       `gorm:"primaryKey;autoIncrement"`
	UserID      uint       `gorm:"uniqueIndex;not null"`
	FullName    string     `gorm:"type:varchar(255)"`
	DateOfBirth *time.Time `gorm:"type:date"`
	Occupation  string     `gorm:"type:varchar(100)"`
	Address     string     `gorm:"type:text"`
	City        string     `gorm:"type:varchar(100)"`
	Province    string     `gorm:"type:varchar(100)"`
	PhotoURL    string     `gorm:"type:varchar(500)"`

	User         User          `gorm:"foreignKey:UserID"`
	Transactions []Transaction `gorm:"foreignKey:ReceiptID"`

	Timestamp
}