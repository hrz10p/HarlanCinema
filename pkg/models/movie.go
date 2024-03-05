package models

type Movie struct {
	ID          int64   `gorm:"primaryKey"`
	Title       string  `gorm:"size:255;not null"`
	Description string  `gorm:"type:text;not null"`
	Rating      float64 `gorm:"not null"`
	ImageUrl    string  `gorm:"not null"`
}
