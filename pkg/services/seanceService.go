package services

import (
	"HarlanCinema/pkg/models"
	"gorm.io/gorm"
)

type SeanceService struct {
	db *gorm.DB
}

func NewSeanceService(db *gorm.DB) *SeanceService {
	return &SeanceService{db: db}
}

func (ss *SeanceService) Create(seance models.Seance) (models.Seance, error) {
	if err := ss.db.Create(&seance).Error; err != nil {
		return models.Seance{}, err
	}
	return seance, nil
}

func (ss *SeanceService) GetAll() ([]models.Seance, error) {
	var seances []models.Seance
	if err := ss.db.Preload("Movie").Find(&seances).Error; err != nil {
		return nil, err
	}
	return seances, nil
}

func (ss *SeanceService) GetByID(id int64) (models.Seance, error) {
	var seance models.Seance
	if err := ss.db.Preload("Movie").First(&seance, id).Error; err != nil {
		return models.Seance{}, err
	}
	return seance, nil
}

func (ss *SeanceService) Update(seance models.Seance) (models.Seance, error) {
	if err := ss.db.Save(&seance).Error; err != nil {
		return models.Seance{}, err
	}
	return seance, nil
}

func (ss *SeanceService) Delete(id int64) error {
	if err := ss.db.Delete(&models.Seance{}, id).Error; err != nil {
		return err
	}
	return nil
}
