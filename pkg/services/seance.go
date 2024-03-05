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

func (ss *SeanceService) CreateSeance(seance models.Seance) (models.Seance, error) {
	seance, err := ss.Repo.SeanceRepository.Create(seance)
	return seance, err
}

func (ss *SeanceService) GetAllSeances() ([]models.Seance, error) {
	seances, err := ss.Repo.SeanceRepository.GetAll()
	return seances, err
}

func (ss *SeanceService) GetSeanceById(seanceId int64) (models.Seance, error) {
	seance, err := ss.Repo.SeanceRepository.GetByID(seanceId)
	return seance, err
}

func (ss *SeanceService) UpdateSeance(seance models.Seance) (models.Seance, error) {
	seance, err := ss.Repo.SeanceRepository.Update(seance)
	return seance, err
}

func (ss *SeanceService) DeleteSeance(seanceId int64) error {
	if err := ss.Repo.SeanceRepository.Delete(seanceId); err != nil {
		return err
	}
	return nil
}
