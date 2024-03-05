package services

import (
	"HarlanCinema/pkg/models"
	repo "HarlanCinema/pkg/repos"
)

type SeanceService struct {
	Repo *repo.Repository
}

func NewSeanceService(Repo *repo.Repository) *SeanceService {
	return &SeanceService{Repo: Repo}
}

func (ss *SeanceService) Create(seance models.Seance) (models.Seance, error) {
	seance, err := ss.Repo.SeanceRepository.Create(seance)
	return seance, err
}
