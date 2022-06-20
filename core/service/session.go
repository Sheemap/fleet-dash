package service

import (
	"errors"
	"fleet-dash-core/data"
	"fleet-dash-core/eveclient"
	"github.com/ReneKroon/ttlcache/v2"
	"github.com/go-kit/kit/log"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"strconv"
	"time"
)

type sessionService struct {
	repo       data.Repository
	fleetCache *ttlcache.Cache
	logger     log.Logger
}

var (
	ErrSessionAlreadyRunning = errors.New("session already running")
	ErrNotInFleet            = errors.New("not in fleet")
	ErrNotInSession          = errors.New("not in active session")
)

type SessionService interface {
	StartSession(character Character) (string, error)
	EndSession(c Character) error
	GetCharacterActiveSession(token *jwt.Token) (*data.Session, error)
}

func NewSessionService(repository data.Repository, logger log.Logger) SessionService {
	fleetCache := ttlcache.NewCache()
	err := fleetCache.SetTTL(time.Second * 60)
	if err != nil {
		panic(err)
	}

	fleetCache.SkipTTLExtensionOnHit(true)

	return &sessionService{
		repo:       repository,
		fleetCache: fleetCache,
		logger:     logger,
	}
}

func (s *sessionService) EndSession(c Character) error {
	currentSession, err := s.GetCharacterActiveSession(c.AccessToken)
	if err != nil {
		return err
	}
	if currentSession == nil {
		return ErrNotInSession
	}

	return s.repo.EndSession(currentSession.ID)
}

func (s *sessionService) StartSession(c Character) (string, error) {
	fleetID, err := s.getCharacterFleet(c.AccessToken)
	if err != nil {
		return "", err
	}

	fleetSession, err := s.repo.GetSessionByFleet(*fleetID)
	if err != nil {
		return "", err
	}
	if fleetSession != nil {
		return "", ErrSessionAlreadyRunning
	}

	newID := uuid.New().String()

	newSession := &data.Session{
		BaseModel: data.BaseModel{
			ID: newID,
		},
		FleetID:     *fleetID,
		CharacterID: c.ID,
	}

	err = s.repo.SaveSession(newSession)

	return newSession.ID, err
}

func (s *sessionService) GetCharacterActiveSession(token *jwt.Token) (*data.Session, error) {
	fleetID, err := s.getCharacterFleet(token)
	if err != nil {
		return nil, err
	}
	if fleetID == nil {
		return nil, ErrNotInFleet
	}

	fleetSession, err := s.repo.GetSessionByFleet(*fleetID)
	if err != nil {
		return nil, err
	}
	if fleetSession == nil {
		return nil, ErrNotInSession
	}

	return fleetSession, nil
}

func (s *sessionService) getCharacterFleet(token *jwt.Token) (*string, error) {
	// hit cache for fleet ID
	strFleetID, err := s.fleetCache.Get(token.Raw)

	// if not found, get from eve client
	if err == ttlcache.ErrNotFound {
		eveClient, err := eveclient.NewEveClient(token)
		if err != nil {
			return nil, err
		}

		fleet, err := eveClient.GetCurrentFleet()
		if err != nil {
			return nil, err
		}

		if fleet == nil {
			// still cache if nil, means we aren't in fleet
			err = s.fleetCache.Set(token.Raw, nil)
			if err != nil {
				return nil, err
			}
			return nil, nil
		}

		// if not nil, cache fleet id as string
		strFleetID = strconv.FormatInt(*fleet.FleetID, 10)
		err = s.fleetCache.Set(token.Raw, strFleetID)
		if err != nil {
			return nil, err
		}
	}
	if strFleetID == nil {
		return nil, ErrNotInFleet
	}

	// cast our cached fleet id to string
	fleetID := strFleetID.(string)
	return &fleetID, nil
}
