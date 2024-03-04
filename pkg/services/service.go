package services

import (
	"gorm.io/gorm"
)

type Service struct {
	UserService    UserService
	SessionService SessionService
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		UserService:    *NewUserService(db),
		SessionService: *NewSessionService(db),
	}
}
