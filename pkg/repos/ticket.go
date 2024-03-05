package repo

import (
	"HarlanCinema/pkg/models"
	"gorm.io/gorm"
)

type TicketRepository struct {
	db *gorm.DB
}

func NewTicketRepository(db *gorm.DB) *TicketRepository {
	return &TicketRepository{db: db}
}

func (ts *TicketRepository) Create(ticket models.Ticket) (models.Ticket, error) {
	if err := ts.db.Create(&ticket).Error; err != nil {
		return models.Ticket{}, err
	}
	return ticket, nil
}

func (ts *TicketRepository) GetAll() ([]models.Ticket, error) {
	var tickets []models.Ticket
	if err := ts.db.Preload("User").Preload("Seance").Find(&tickets).Error; err != nil {
		return nil, err
	}
	return tickets, nil
}

func (ts *TicketRepository) GetByID(userID, seanceID int64) (models.Ticket, error) {
	var ticket models.Ticket
	if err := ts.db.Preload("User").Preload("Seance").Where("user_id = ? AND seance_id = ?", userID, seanceID).First(&ticket).Error; err != nil {
		return models.Ticket{}, err
	}
	return ticket, nil
}

func (ts *TicketRepository) Update(ticket models.Ticket) (models.Ticket, error) {
	if err := ts.db.Save(&ticket).Error; err != nil {
		return models.Ticket{}, err
	}
	return ticket, nil
}

func (ts *TicketRepository) Delete(userID, seanceID int64) error {
	if err := ts.db.Delete(&models.Ticket{}, "user_id = ? AND seance_id = ?", userID, seanceID).Error; err != nil {
		return err
	}
	return nil
}

func (ts *TicketRepository) FindAllTicketsByUserID(userID string) ([]models.Ticket, error) {
	var tickets []models.Ticket
	if err := ts.db.Preload("User").Preload("Seance").Where("user_id = ?", userID).Find(&tickets).Error; err != nil {
		return nil, err
	}
	return tickets, nil
}
