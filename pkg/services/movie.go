package services

import (
	"HarlanCinema/pkg/models"
	repo "HarlanCinema/pkg/repos"
)

type MovieService struct {
	Repo *repo.Repository
}

func NewMovieService(Repo *repo.Repository) *MovieService {
	return &MovieService{Repo: Repo}
}

func (ms *MovieService) Create(movie models.Movie) (models.Movie, error) {
	m, err := ms.Repo.MovieRepository.Create(movie)
	return m, err
}
