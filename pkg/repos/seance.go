package repositories

import (
	"HarlanCinema/pkg/models"
	"gorm.io/gorm"
)

type SeanceRepository struct {
	db *gorm.DB
}

func NewSeanceRepository(db *gorm.DB) *SeanceRepository {
	return &SeanceRepository{db: db}
}

func (ss *SeanceRepository) Create(seance models.Seance) (models.Seance, error) {
	if err := ss.db.Create(&seance).Error; err != nil {
		return models.Seance{}, err
	}
	return seance, nil
}

func (ss *SeanceRepository) GetAll() ([]models.Seance, error) {
	var seances []models.Seance
	if err := ss.db.Preload("Movie").Find(&seances).Error; err != nil {
		return nil, err
	}
	return seances, nil
}

func (ss *SeanceRepository) GetByID(id int64) (models.Seance, error) {
	var seance models.Seance
	if err := ss.db.Preload("Movie").First(&seance, id).Error; err != nil {
		return models.Seance{}, err
	}
	return seance, nil
}

func (ss *SeanceRepository) Update(seance models.Seance) (models.Seance, error) {
	if err := ss.db.Save(&seance).Error; err != nil {
		return models.Seance{}, err
	}
	return seance, nil
}

func (ss *SeanceRepository) Delete(id int64) error {
	if err := ss.db.Delete(&models.Seance{}, id).Error; err != nil {
		return err
	}
	return nil
}
