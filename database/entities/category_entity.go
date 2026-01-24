package entities

type Category struct {
	ID   uint   `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"type:varchar(100);not null;uniqueIndex"`

	SurplusListings []SurplusListing `gorm:"foreignKey:CategoryID"`

	Timestamp
}
