package entities

import (
	"time"
)

type FoodCondition string

const (
	FoodConditionExcellent FoodCondition = "excellent"
	FoodConditionGood      FoodCondition = "good"
	FoodConditionFair      FoodCondition = "fair"
)

type SurplusStatus string

const (
	SurplusStatusAvailable SurplusStatus = "available"
	SurplusStatusSoldOut   SurplusStatus = "sold_out"
	SurplusStatusExpired   SurplusStatus = "expired"
)

type SurplusListing struct {
	ID                 uint          `gorm:"primaryKey;autoIncrement"`
	MitraID            uint          `gorm:"not null;index" json:"mitra_id"`
	CategoryID         *uint         `gorm:"index" json:"category_id"`
	Name               string        `gorm:"type:varchar(255);not null" json:"name"`
	Description        string        `gorm:"type:text" json:"description"`
	ImageURL           string        `gorm:"type:varchar(500)" json:"image_url"`
	OriginalPrice      int           `gorm:"not null" json:"original_price"`
	DiscountedPrice    int           `gorm:"not null" json:"discounted_price"`
	DiscountPercentage int           `gorm:"default:0" json:"discount_percentage"`
	PortionsSold       int           `gorm:"default:0" json:"portions_sold"`
	PortionsAvailable  int           `gorm:"not null" json:"portions_available"`
	FoodCondition      FoodCondition `gorm:"type:varchar(20);not null;default:'good'" json:"food_condition"`
	Status             SurplusStatus `gorm:"type:varchar(20);not null;default:'available'" json:"status"`
	PreparedAt         *time.Time    `gorm:"type:timestamp with time zone" json:"prepared_at"`
	ExpiresAt          *time.Time    `gorm:"type:timestamp with time zone;not null" json:"expires_at"`

	Mitra            MitraProfile      `gorm:"foreignKey:MitraID" json:"mitra,omitempty"`
	Category         *Category         `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	TransactionItems []TransactionItem `gorm:"foreignKey:SurplusListingID" json:"transaction_items,omitempty"`

	Timestamp
}

