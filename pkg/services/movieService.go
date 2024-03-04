package services

import (
	"HarlanCinema/pkg/models"
	"gorm.io/gorm"
)

type MovieService struct {
	db *gorm.DB
}

func NewMovieService(db *gorm.DB) *MovieService {
	return &MovieService{db: db}
}

func (ms *MovieService) Create(movie models.Movie) (models.Movie, error) {
	if err := ms.db.Create(&movie).Error; err != nil {
		return models.Movie{}, err
	}
	return movie, nil
}

func (ms *MovieService) GetAll() ([]models.Movie, error) {
	var movies []models.Movie
	if err := ms.db.Find(&movies).Error; err != nil {
		return nil, err
	}
	return movies, nil
}

func (ms *MovieService) GetByID(id int64) (models.Movie, error) {
	var movie models.Movie
	if err := ms.db.First(&movie, id).Error; err != nil {
		return models.Movie{}, err
	}
	return movie, nil
}

func (ms *MovieService) Update(movie models.Movie) (models.Movie, error) {
	if err := ms.db.Save(&movie).Error; err != nil {
		return models.Movie{}, err
	}
	return movie, nil
}

func (ms *MovieService) Delete(id int64) error {
	if err := ms.db.Delete(&models.Movie{}, id).Error; err != nil {
		return err
	}
	return nil
}
