package models

import "time"

type Seance struct {
	ID       int64     `gorm:"primaryKey"`
	MovieID  int64     `gorm:"not null"`
	Date     time.Time `gorm:"not null"`
	Location string    `gorm:"size:255;not null"`
	Movie    Movie     `gorm:"foreignKey:MovieID"`
}
