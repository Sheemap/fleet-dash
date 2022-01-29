package service

import (
	"errors"
	"fleet-dash-core/data"
	"fleet-dash-core/eveclient"
	"fleet-dash-core/utilities"
	"github.com/ReneKroon/ttlcache/v2"
	"github.com/go-kit/kit/log"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"strconv"
	"time"
)

type sessionService struct {
	repo data.Repository
	fleetCache *ttlcache.Cache
	logger log.Logger
}

var (
	ErrSessionAlreadyRunning = errors.New("session already running")
	ErrNotInFleet			= errors.New("not in fleet")
	ErrNotInSession = errors.New("not in active session")
)

type SessionService interface {
	StartSession(character Character) (string, error)
	GetCharacterActiveSession(token *jwt.Token) (*string, error)
}

func NewSessionService(repository data.Repository, logger log.Logger) SessionService {
	fleetCache := ttlcache.NewCache()
	err := fleetCache.SetTTL(time.Second * 60)
	if err != nil {
		panic(err)
	}

	fleetCache.SkipTTLExtensionOnHit(true)

	return &sessionService{
		repo: repository,
		fleetCache: fleetCache,
		logger: logger,
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

	// hit cache for fleet ID
	strFleetID, err := s.fleetCache.Get(token.Raw)

	// if not found, get from eve client
	if err == ttlcache.ErrNotFound {
		err = s.logger.Log("msg", "cache miss, fetching fleet from API", "characterID", charID)
		if err != nil {
			return nil, err
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

	fleetSession, err := s.repo.GetSessionByFleet(fleetID)
	if err != nil {
		return nil, err
	}
	if fleetSession == nil {
		return nil, ErrNotInSession
	}

	return &fleetSession.ID, nil
}
