package service

import (
	"fleet-dash-core/data"
	"github.com/google/uuid"
	"time"
)

type eventService struct{
	repo data.Repository
}

type EventService interface {
	PersistEveLogEvent(sessionID string, event EveLogEvent) error
}

type EveLogEvent struct {
	Timestamp   time.Time
	Type        string
	CharacterID string
	Amount int32
	Pilot string
	Ship string
	Weapon string
	Application string
	Corporation string
	Alliance string
}

func NewEventService(r data.Repository) EventService {
	return &eventService{
		repo: r,
	}
}

func (s *eventService) PersistEveLogEvent(sessionID string, e EveLogEvent) error {
	event := data.Event{
		BaseModel: data.BaseModel{ID: uuid.New().String()},
		SessionID: sessionID,
		Timestamp: e.Timestamp,
		Type: e.Type,
		CharacterID: e.CharacterID,
		Amount: e.Amount,
		Pilot: e.Pilot,
		Ship: e.Ship,
		Weapon: e.Weapon,
		Application: e.Application,
		Corporation: e.Corporation,
		Alliance: e.Alliance,
	}

	return s.repo.SaveEveLogEvent(&event)
}

