package service

import (
	"errors"
	"fleet-dash-core/data"
	"fleet-dash-core/eveclient"
	"fleet-dash-core/utilities"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"strconv"
)

type sessionService struct {
	repo data.Repository
}

var (
	ErrSessionAlreadyRunning = errors.New("session already running")
	ErrNotInFleet			= errors.New("not in fleet")
)

type SessionService interface {
	StartSession(character Character) (string, error)
	GetCharacterActiveSession(token *jwt.Token) (*string, error)
}

func NewSessionService(repository data.Repository) SessionService {
	return &sessionService{
		repo: repository,
	}
}

func (s *sessionService) StartSession(c Character) (string, error) {
	currentSession, err := s.repo.GetCharacterActiveSession(c.ID)
	if err != nil {
		return "", err
	}
	if currentSession != nil {
		return "", ErrSessionAlreadyRunning
	}

	eveClient, err := eveclient.NewEveClient(c.AccessToken)
	if err != nil {
		return "", err
	}

	fleet, err := eveClient.GetCurrentFleet()
	if err != nil {
		return "", err
	}
	if fleet == nil {
		return "", ErrNotInFleet
	}
	strFleetID := strconv.FormatInt(*fleet.FleetID, 10)
	fleetSession, err := s.repo.GetSessionByFleet(strFleetID)
	if err != nil {
		return "", err
	}
	if fleetSession != nil {
		return "", ErrSessionAlreadyRunning
	}

	newID := uuid.New().String()

	newSession := &data.Session{
		BaseModel: data.BaseModel{
			ID:	newID,
		},
		FleetID: strFleetID,
		CharacterID: c.ID,
	}

	err = s.repo.SaveSession(newSession)

	return newSession.ID, err
}

func (s *sessionService) GetCharacterActiveSession(token *jwt.Token) (*string, error) {
	charID, err := utilities.GetCharacterIDFromToken(token)
	if err != nil {
		return nil, err
	}

	session, err := s.repo.GetCharacterActiveSession(charID)
	if err != nil {
		return nil, err
	}
	if session != nil {
		return &session.ID, nil
	}

	eveClient, err := eveclient.NewEveClient(token)
	if err != nil {
		return nil, err
	}

	fleet, err := eveClient.GetCurrentFleet()
	if err != nil {
		return nil, err
	}
	if fleet == nil {
		return nil, ErrNotInFleet
	}
	strFleetID := strconv.FormatInt(*fleet.FleetID, 10)
	fleetSession, err := s.repo.GetSessionByFleet(strFleetID)
	if err != nil {
		return nil, err
	}

	return &fleetSession.ID, nil
}
