package entities

type TransactionItem struct {
	ID               uint   `gorm:"primaryKey;autoIncrement"`
	TransactionID    uint   `gorm:"not null;index"`
	SurplusListingID uint   `gorm:"not null;index"`
	Quantity         int    `gorm:"not null"`
	Subtotal         int    `gorm:"not null"`
	Notes            string `gorm:"type:text"`

	Transaction    Transaction    `gorm:"foreignKey:TransactionID"`
	SurplusListing SurplusListing `gorm:"foreignKey:SurplusListingID"`

	Timestamp
}