package service

import "fleet-dash-core/data"

type eventStreamService struct{
	repo data.Repository
}

type EventStreamService interface {
	GenerateEventStreamTicket(sessionId string) (*string, error)
}

func NewEventStreamService(r data.Repository) EventStreamService {
	return &eventStreamService{
		repo: r,
	}
}

func (e *eventStreamService)GenerateEventStreamTicket(sessionId string) (*string, error){
	ticket, err := e.repo.GenerateEventStreamTicket(sessionId)
	if err != nil {
		return nil, err
	}

	return &ticket.Ticket, nil
}