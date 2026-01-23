package entities

type TransactionItem struct {
	ID               uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	TransactionID    uint   `gorm:"not null;index" json:"transaction_id"`
	SurplusListingID uint   `gorm:"not null;index" json:"surplus_listing_id"`
	Quantity         int    `gorm:"not null" json:"quantity"`
	Subtotal         int    `gorm:"not null" json:"subtotal"`
	Notes            string `gorm:"type:text" json:"notes"`

	Transaction    Transaction    `gorm:"foreignKey:TransactionID" json:"transaction,omitempty"`
	SurplusListing SurplusListing `gorm:"foreignKey:SurplusListingID" json:"surplus_listing,omitempty"`

	Timestamp
}