package services

import (
	"HarlanCinema/pkg/models"
	repo "HarlanCinema/pkg/repos"
)

type ReviewService struct {
	Repo *repo.Repository
}

func NewReviewService(Repo *repo.Repository) *ReviewService {
	return &ReviewService{Repo: Repo}
}

func (rs *ReviewService) CreateReview(review models.Review) (models.Review, error) {
	rev, err := rs.Repo.ReviewRepository.Create(review)
	return rev, err
}

func (rs *ReviewService) GetAllReviews() ([]models.Review, error) {
	reviews, err := rs.Repo.ReviewRepository.GetAll()
	return reviews, err
}

func (rs *ReviewService) GetReviewByUserAndMovieId(userId, movieId int64) (models.Review, error) {
	review, err := rs.Repo.ReviewRepository.GetByUserAndMovieID(userId, movieId)
	return review, err
}

func (rs *ReviewService) UpdateReview(review models.Review) (models.Review, error) {
	review, err := rs.Repo.ReviewRepository.Update(review)
	return review, err
}

func (rs *ReviewService) DeleteReview(userId, movieId int64) error {
	if err := rs.Repo.ReviewRepository.Delete(userId, movieId); err != nil {
		return err
	}
	return nil
}
