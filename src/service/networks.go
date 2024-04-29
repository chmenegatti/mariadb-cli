package service

import (
	"context"

	"nemesis-cli/src/models"
	"nemesis-cli/src/repository"
)

type NetworkService interface {
	GetAll(ctx context.Context) ([]models.Networks, error)
	GetByID(ctx context.Context, id string) (models.Networks, error)
	GetByName(ctx context.Context, name string) ([]models.Networks, error)
}

type networkService struct {
	repository repository.NetworkRepository
}

func NewNetworkService(repository repository.NetworkRepository) NetworkService {
	return &networkService{
		repository: repository,
	}
}

func (s *networkService) GetAll(ctx context.Context) ([]models.Networks, error) {
	return s.repository.GetAll(ctx)
}

func (s *networkService) GetByID(ctx context.Context, id string) (models.Networks, error) {
	return s.repository.GetByID(ctx, id)
}

func (s *networkService) GetByName(ctx context.Context, name string) ([]models.Networks, error) {
	return s.repository.FindByName(ctx, name)
}
