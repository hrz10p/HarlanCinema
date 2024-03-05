package services

import (
	"HarlanCinema/pkg/models"
	repo "HarlanCinema/pkg/repos"
)

type TicketService struct {
	Repo *repo.Repository
}

func NewTicketService(Repo *repo.Repository) *TicketService {
	return &TicketService{Repo: Repo}
}

func (ts *TicketService) CreateTicket(ticket models.Ticket) (models.Ticket, error) {
	tick, err := ts.Repo.TicketRepository.Create(ticket)
	return tick, err
}

func (ts *TicketService) GetAllTickets() ([]models.Ticket, error) {
	tickets, err := ts.Repo.TicketRepository.GetAll()
	return tickets, err
}

func (ts *TicketService) GetTicketById(userId, seanceId int64) (models.Ticket, error) {
	ticket, err := ts.Repo.TicketRepository.GetByID(userId, seanceId)
	return ticket, err
}

func (ts *TicketService) UpdateTicket(ticket models.Ticket) (models.Ticket, error) {
	ticket, err := ts.Repo.TicketRepository.Update(ticket)
	return ticket, err
}

func (ts *TicketService) DeleteTicket(userId, seanceId int64) error {
	if err := ts.Repo.TicketRepository.Delete(userId, seanceId); err != nil {
		return err
	}
	return nil
}

func (ts *TicketService) GiveTicketForUser(userID string, seanceID int64) error {
	ticket := models.Ticket{
		UserID:     userID,
		SeanceID:   seanceID,
		Cost:       0,
		TicketType: "plain",
	}
	if _, err := ts.Repo.TicketRepository.Create(ticket); err != nil {
		return err
	}
	return nil
}
