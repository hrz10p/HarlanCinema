package repo

import (
	"HarlanCinema/pkg/models"
	"gorm.io/gorm"
)

type MovieRepository struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) *MovieRepository {
	return &MovieRepository{db: db}
}

func (ms *MovieRepository) Create(movie models.Movie) (models.Movie, error) {
	if err := ms.db.Create(&movie).Error; err != nil {
		return models.Movie{}, err
	}
	return movie, nil
}

func (ms *MovieRepository) GetAll() ([]models.Movie, error) {
	var movies []models.Movie
	if err := ms.db.Find(&movies).Error; err != nil {
		return nil, err
	}
	return movies, nil
}

func (ms *MovieRepository) GetByID(id int64) (models.Movie, error) {
	var movie models.Movie
	if err := ms.db.First(&movie, id).Error; err != nil {
		return models.Movie{}, err
	}
	return movie, nil
}

func (ms *MovieRepository) Update(movie models.Movie) (models.Movie, error) {
	if err := ms.db.Save(&movie).Error; err != nil {
		return models.Movie{}, err
	}
	return movie, nil
}

func (ms *MovieRepository) Delete(id int64) error {
	if err := ms.db.Delete(&models.Movie{}, id).Error; err != nil {
		return err
	}
	return nil
}
