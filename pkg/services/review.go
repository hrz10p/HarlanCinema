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

func (rs *ReviewService) Create(review models.Review) (models.Review, error) {
	rev, err := rs.Repo.ReviewRepository.Create(review)
	return rev, err
}
