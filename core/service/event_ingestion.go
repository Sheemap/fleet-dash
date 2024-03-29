package service

import (
	"fleet-dash-core/data"
	"github.com/google/uuid"
	"time"
)

type eventIngestionService struct {
	repo data.Repository
}

type EventIngestionService interface {
	PersistEveLogEvent(sessionID string, event EveLogEvent) error
	PersistEveLogEventBatch(sessionID string, eventBatch []EveLogEvent) error
}

type EveLogEvent struct {
	Timestamp           time.Time
	Type                string
	CharacterID         string
	CharacterShipTypeID int32
	Amount              int32
	Pilot               string
	Ship                string
	Weapon              string
	Application         string
	Corporation         string
	Alliance            string
}

func NewEventIngestionService(r data.Repository) EventIngestionService {
	return &eventIngestionService{
		repo: r,
	}
}

func (s *eventIngestionService) PersistEveLogEvent(sessionID string, e EveLogEvent) error {
	event := data.Event{
		BaseModel:           data.BaseModel{ID: uuid.New().String()},
		SessionID:           sessionID,
		Timestamp:           e.Timestamp,
		Type:                e.Type,
		CharacterID:         e.CharacterID,
		CharacterShipTypeID: e.CharacterShipTypeID,
		Amount:              e.Amount,
		Pilot:               e.Pilot,
		Ship:                e.Ship,
		Weapon:              e.Weapon,
		Application:         e.Application,
		Corporation:         e.Corporation,
		Alliance:            e.Alliance,
	}

	return s.repo.SaveEveLogEvent(&event)
}

func (s *eventIngestionService) PersistEveLogEventBatch(sessionID string, eventBatch []EveLogEvent) error {
	events := make([]data.Event, len(eventBatch))

	for i, e := range eventBatch {
		events[i] = data.Event{
			BaseModel:           data.BaseModel{ID: uuid.New().String()},
			SessionID:           sessionID,
			Timestamp:           e.Timestamp,
			Type:                e.Type,
			CharacterID:         e.CharacterID,
			CharacterShipTypeID: e.CharacterShipTypeID,
			Amount:              e.Amount,
			Pilot:               e.Pilot,
			Ship:                e.Ship,
			Weapon:              e.Weapon,
			Application:         e.Application,
			Corporation:         e.Corporation,
			Alliance:            e.Alliance,
		}
	}

	return s.repo.SaveEveLogEventBatch(&events)
}
