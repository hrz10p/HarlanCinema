package models

import (
	"time"
)

type Session struct {
	ID         string    `gorm:"type:uuid;primary_key;"`
	UID        string    `gorm:"not null;"`
	ExpireTime time.Time `gorm:"not null;"`
}
