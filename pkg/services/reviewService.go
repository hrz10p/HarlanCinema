package services

import (
	"HarlanCinema/pkg/models"
	"gorm.io/gorm"
)

type ReviewService struct {
	db *gorm.DB
}

func NewReviewService(db *gorm.DB) *ReviewService {
	return &ReviewService{db: db}
}

func (rs *ReviewService) Create(review models.Review) (models.Review, error) {
	if err := rs.db.Create(&review).Error; err != nil {
		return models.Review{}, err
	}
	return review, nil
}

func (rs *ReviewService) GetAll() ([]models.Review, error) {
	var reviews []models.Review
	if err := rs.db.Preload("User").Preload("Movie").Find(&reviews).Error; err != nil {
		return nil, err
	}
	return reviews, nil
}

func (rs *ReviewService) GetByUserAndMovieID(userID, movieID int64) (models.Review, error) {
	var review models.Review
	if err := rs.db.Where("user_id = ? AND movie_id = ?", userID, movieID).Preload("User").Preload("Movie").First(&review).Error; err != nil {
		return models.Review{}, err
	}
	return review, nil
}

func (rs *ReviewService) Update(review models.Review) (models.Review, error) {
	if err := rs.db.Model(&models.Review{}).Where("user_id = ? AND movie_id = ?", review.UserID, review.MovieID).Updates(review).Error; err != nil {
		return models.Review{}, err
	}
	return review, nil
}

func (rs *ReviewService) Delete(userID, movieID int64) error {
	if err := rs.db.Where("user_id = ? AND movie_id = ?", userID, movieID).Delete(&models.Review{}).Error; err != nil {
		return err
	}
	return nil
}
