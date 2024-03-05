package services

import (
	repo "HarlanCinema/pkg/repos"
)

type Service struct {
	UserService   UserService
	MovieService  MovieService
	TicketService TicketService
	ReviewService ReviewService
	SeanceService SeanceService
}

func NewService(Repo *repo.Repository) *Service {
	return &Service{
		UserService:   *NewUserService(Repo),
		MovieService:  *NewMovieService(Repo),
		TicketService: *NewTicketService(Repo),
		ReviewService: *NewReviewService(Repo),
		SeanceService: *NewSeanceService(Repo),
	}
}
