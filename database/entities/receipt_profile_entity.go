package entities

import (
	"time"
)

type ReceiptProfile struct {
	ID          uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID      uint       `gorm:"uniqueIndex;not null" json:"user_id"`
	FullName    string     `gorm:"type:varchar(255)" json:"full_name"`
	DateOfBirth *time.Time `gorm:"type:date" json:"date_of_birth"`
	Occupation  string     `gorm:"type:varchar(100)" json:"occupation"`
	Address     string     `gorm:"type:text" json:"address"`
	City        string     `gorm:"type:varchar(100)" json:"city"`
	Province    string     `gorm:"type:varchar(100)" json:"province"`
	PhotoURL    string     `gorm:"type:varchar(500)" json:"photo_url"`

	User         User          `gorm:"foreignKey:UserID" json:"-"`
	Transactions []Transaction `gorm:"foreignKey:ReceiptID" json:"transactions,omitempty"`

	Timestamp
}