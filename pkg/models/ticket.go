package models

type Ticket struct {
	UserID     string `gorm:"primaryKey"`
	SeanceID   int64  `gorm:"primaryKey"`
	Cost       int64  `gorm:"not null"`
	TicketType string `gorm:"size:255;not null"`
	User       User   `gorm:"foreignKey:UserID"`
	Seance     Seance `gorm:"foreignKey:SeanceID"`
}
