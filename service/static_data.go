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
	GetStaticItem(id uint64) (*data.StaticItemInfo, error)
	GetStaticSolarSystem(id uint64) (*data.StaticSolarSystemInfo, error)
}

func NewStaticDataService(repository data.Repository, logger log.Logger) StaticDataService {
	return &staticDataService{
		repo:   repository,
		logger: logger,
	}
}

func (svc *staticDataService) GetStaticItem(id uint64) (*data.StaticItemInfo, error) {
	item, err := svc.repo.GetStaticItemInfo(id)
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
