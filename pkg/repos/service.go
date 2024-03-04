package repositories

import (
	"gorm.io/gorm"
)

type Repository struct {
	UserRepository    UserRepository
	SessionRepository SessionRepository
	MovieRepository   MovieRepository
	TicketRepository  TicketRepository
	ReviewRepository  ReviewRepository
	SeanceRepository  SeanceRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		UserRepository:    *NewUserRepository(db),
		SessionRepository: *NewSessionRepository(db),
		MovieRepository:   *NewMovieRepository(db),
		TicketRepository:  *NewTicketRepository(db),
		ReviewRepository:  *NewReviewRepository(db),
		SeanceRepository:  *NewSeanceRepository(db),
	}
}
