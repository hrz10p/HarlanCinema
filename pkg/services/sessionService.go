package services

import (
	"HarlanCinema/pkg/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SessionService struct {
	db *gorm.DB
}

func NewSessionService(db *gorm.DB) *SessionService {
	return &SessionService{db: db}
}

// RegisterSession creates or updates a session for a given UID with an expiration time.
// It uses a transaction to ensure that these operations are atomic.
func (s *SessionService) RegisterSession(UID string, exp time.Time) (models.Session, error) {
	var session models.Session
	err := s.db.Transaction(func(tx *gorm.DB) error {
		// Check if a session exists for the UID and delete it if found
		if err := tx.Where("uid = ?", UID).First(&session).Error; err == nil {
			if err := tx.Delete(&models.Session{}, session.ID).Error; err != nil {
				return err
			}
		} else if err != gorm.ErrRecordNotFound {
			return err
		}

		// Create a new session
		session = models.Session{ID: uuid.New().String(), UID: UID, ExpireTime: exp}
		if err := tx.Create(&session).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return models.Session{}, err
	}

	return session, nil
}
