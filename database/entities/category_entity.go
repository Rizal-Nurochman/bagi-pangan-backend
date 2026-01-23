package entities

type Category struct {
	ID   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"type:varchar(100);not null;uniqueIndex" json:"name"`

	SurplusListings []SurplusListing `gorm:"foreignKey:CategoryID" json:"surplus_listings,omitempty"`

	Timestamp
}
