package models

type User struct {
	ID       string `gorm:"type:uuid;primary_key;"`
	Username string `gorm:"not null;unique;"`
	Email    string `gorm:"not null;unique;"`
	Password string `gorm:"not null;"`
	Role     string `gorm:"not null;"`
}
