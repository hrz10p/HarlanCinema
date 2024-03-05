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

func (ts *TicketService) Create(ticket models.Ticket) (models.Ticket, error) {
	tick, err := ts.Repo.TicketRepository.Create(ticket)
	return tick, err
}
