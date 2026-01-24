package entities

type Review_Mitra struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	UserID    uint      `gorm:"not null;index"`
	MitraID   uint      `gorm:"not null;index"`
	Rating    float32   `gorm:"not null"`
	Comment   string    `gorm:"type:text"`
	ImageURL	string    `gorm:"type:varchar(500)"`
	User      User      `gorm:"foreignKey:UserID"`
	Mitra     MitraProfile `gorm:"foreignKey:MitraID"`
	Timestamp
}