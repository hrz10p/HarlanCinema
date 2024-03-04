package models

type Review struct {
	UserID  int64   `gorm:"primaryKey"`
	MovieID int64   `gorm:"primaryKey"`
	Text    string  `gorm:"type:text;not null"`
	Rate    float64 `gorm:"not null"`
	User    User    `gorm:"foreignKey:UserID"`
	Movie   Movie   `gorm:"foreignKey:MovieID"`
}
