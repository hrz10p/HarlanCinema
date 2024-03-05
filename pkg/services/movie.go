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

func (ms *MovieService) CreateMovie(movie models.Movie) (models.Movie, error) {
	m, err := ms.Repo.MovieRepository.Create(movie)
	return m, err
}

func (ms *MovieService) GetAllMovies() ([]models.Movie, error) {
	movies, err := ms.Repo.MovieRepository.GetAll()
	return movies, err
}

func (ms *MovieService) GetMovieById(movieId int64) (models.Movie, error) {
	movie, err := ms.Repo.MovieRepository.GetByID(movieId)
	return movie, err
}

func (ms *MovieService) UpdateMovie(movie models.Movie) (models.Movie, error) {
	movie, err := ms.Repo.MovieRepository.Update(movie)
	return movie, err
}

func (ms *MovieService) DeleteMovieById(movieId int64) error {
	if err := ms.Repo.MovieRepository.Delete(movieId); err != nil {
		return err
	}
	return nil
}
