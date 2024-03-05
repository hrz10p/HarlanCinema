package repo

import (
	"HarlanCinema/pkg/models"
	"gorm.io/gorm"
)

type ReviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) *ReviewRepository {
	return &ReviewRepository{db: db}
}

func (rs *ReviewRepository) Create(review models.Review) (models.Review, error) {
	if err := rs.db.Create(&review).Error; err != nil {
		return models.Review{}, err
	}
	return review, nil
}

func (rs *ReviewRepository) GetAll() ([]models.Review, error) {
	var reviews []models.Review
	if err := rs.db.Preload("User").Preload("Movie").Find(&reviews).Error; err != nil {
		return nil, err
	}
	return reviews, nil
}

func (rs *ReviewRepository) GetByUserAndMovieID(userID, movieID int64) (models.Review, error) {
	var review models.Review
	if err := rs.db.Where("user_id = ? AND movie_id = ?", userID, movieID).Preload("User").Preload("Movie").First(&review).Error; err != nil {
		return models.Review{}, err
	}
	return review, nil
}

func (rs *ReviewRepository) Update(review models.Review) (models.Review, error) {
	if err := rs.db.Model(&models.Review{}).Where("user_id = ? AND movie_id = ?", review.UserID, review.MovieID).Updates(review).Error; err != nil {
		return models.Review{}, err
	}
	return review, nil
}

func (rs *ReviewRepository) Delete(userID, movieID int64) error {
	if err := rs.db.Where("user_id = ? AND movie_id = ?", userID, movieID).Delete(&models.Review{}).Error; err != nil {
		return err
	}
	return nil
}
