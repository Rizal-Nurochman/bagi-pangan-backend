package entities

import (
	"github.com/shopspring/decimal"
)

type OperationalHours map[string]DaySchedule

type DaySchedule struct {
	Open   string `json:"open"`
	Close  string `json:"close"`
	Closed bool   `json:"closed"`
}

type MitraProfile struct {
	ID                         uint             `gorm:"primaryKey;autoIncrement"`
	UserID                     uint             `gorm:"uniqueIndex;not null"`
	BusinessName               string           `gorm:"type:varchar(255);not null"`
	BusinessType               string           `gorm:"type:varchar(100)"`
	OwnerName                  string           `gorm:"type:varchar(255)"`
	BusinessRegistrationNumber string           `gorm:"type:varchar(100)"`
	PickupInstructions         string           `gorm:"type:varchar(500)"`
	Address                    string           `gorm:"type:text"`
	City                       string           `gorm:"type:varchar(100)"`
	Province                   string           `gorm:"type:varchar(100)"`
	OperationalHours           OperationalHours `gorm:"type:jsonb"`
	BusinessRegistrationDoc    string           `gorm:"type:varchar(500)"`
	OwnerIDCardDoc             string           `gorm:"type:varchar(500)"`
	BusinessPhotoURL           string           `gorm:"type:varchar(500)"`
	TotalPortionsSells         int              `gorm:"default:0"`
	RatingAvg                  decimal.Decimal  `gorm:"type:decimal(3,2);default:0"`
	RatingCount                int              `gorm:"default:0"`

	User            User             `gorm:"foreignKey:UserID"`
	SurplusListings []SurplusListing `gorm:"foreignKey:MitraID"`
	Transactions    []Transaction    `gorm:"foreignKey:MitraID"`

	Timestamp
}