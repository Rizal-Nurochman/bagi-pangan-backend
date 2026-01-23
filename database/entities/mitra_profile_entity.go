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
	ID                         uint             `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID                     uint             `gorm:"uniqueIndex;not null" json:"user_id"`
	BusinessName               string           `gorm:"type:varchar(255);not null" json:"business_name"`
	BusinessType               string           `gorm:"type:varchar(100)" json:"business_type"`
	OwnerName                  string           `gorm:"type:varchar(255)" json:"owner_name"`
	BusinessRegistrationNumber string           `gorm:"type:varchar(100)" json:"business_registration_number"`
	PickupInstructions         string           `gorm:"type:varchar(500)" json:"pickup_instructions"`
	Address                    string           `gorm:"type:text" json:"address"`
	City                       string           `gorm:"type:varchar(100)" json:"city"`
	Province                   string           `gorm:"type:varchar(100)" json:"province"`
	OperationalHours           OperationalHours `gorm:"type:jsonb" json:"operational_hours"`
	BusinessRegistrationDoc    string           `gorm:"type:varchar(500)" json:"business_registration_doc"`
	OwnerIDCardDoc             string           `gorm:"type:varchar(500)" json:"-"`
	BusinessPhotoURL           string           `gorm:"type:varchar(500)" json:"business_photo_url"`
	TotalPortionsSells         int              `gorm:"default:0" json:"total_portions_sells"`
	RatingAvg                  decimal.Decimal  `gorm:"type:decimal(3,2);default:0" json:"rating_avg"`
	RatingCount                int              `gorm:"default:0" json:"rating_count"`

	User            User             `gorm:"foreignKey:UserID" json:"-"`
	SurplusListings []SurplusListing `gorm:"foreignKey:MitraID" json:"surplus_listings,omitempty"`
	Transactions    []Transaction    `gorm:"foreignKey:MitraID" json:"transactions,omitempty"`

	Timestamp
}