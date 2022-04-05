package data

import (
	"time"
)

type BaseModel struct {
	ID        string `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Session struct {
	BaseModel

	EndedAt     *time.Time
	CharacterID string
	FleetID     string

	Events        []Event
	StreamTickets []EventStreamTicket
}

type Event struct {
	BaseModel

	SessionID           string
	Type                string
	Timestamp           time.Time
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

type EventStreamTicket struct {
	BaseModel

	SessionID string
}

type StaticItemInfo struct {
	ID   uint   `gorm:"primarykey"`
	Name string `gorm:"index"`
}

type StaticSolarSystemInfo struct {
	ID   uint `gorm:"primarykey"`
	Name string
}
