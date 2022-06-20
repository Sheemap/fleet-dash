package service

import (
	"fleet-dash-core/data"
	"github.com/go-kit/kit/log"
)

type staticDataService struct {
	repo   data.Repository
	logger log.Logger
}

type StaticDataService interface {
	GetStaticItemByID(id uint64) (*data.StaticItemInfo, error)
	GetStaticSolarSystem(id uint64) (*data.StaticSolarSystemInfo, error)
	GetStaticItemByName(name string) (*data.StaticItemInfo, error)
}

func NewStaticDataService(repository data.Repository, logger log.Logger) StaticDataService {
	return &staticDataService{
		repo:   repository,
		logger: logger,
	}
}

func (svc *staticDataService) GetStaticItemByID(id uint64) (*data.StaticItemInfo, error) {
	item, err := svc.repo.GetStaticItemInfoByID(id)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (svc *staticDataService) GetStaticItemByName(name string) (*data.StaticItemInfo, error) {
	item, err := svc.repo.GetStaticItemInfoByName(name)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (svc *staticDataService) GetStaticSolarSystem(id uint64) (*data.StaticSolarSystemInfo, error) {
	item, err := svc.repo.GetStaticSolarSystemInfo(id)
	if err != nil {
		return nil, err
	}

	return item, nil
}
