package service

import (
	"context"

	"nemesis-cli/src/models"
	"nemesis-cli/src/repository"
)

type OrganizationService interface {
	GetAll(ctx context.Context) ([]models.Organization, error)
	GetByID(ctx context.Context, id string) (models.Organization, error)
	GetByName(ctx context.Context, name string) ([]models.Organization, error)
}

type organizationService struct {
	repository repository.OrganizationRepository
}

func NewOrganizationService(repository repository.OrganizationRepository) OrganizationService {
	return &organizationService{
		repository: repository,
	}
}

func (s *organizationService) GetAll(ctx context.Context) ([]models.Organization, error) {
	return s.repository.GetAll(ctx)
}

func (s *organizationService) GetByID(ctx context.Context, id string) (models.Organization, error) {
	return s.repository.GetByID(ctx, id)
}

func (s *organizationService) GetByName(ctx context.Context, name string) ([]models.Organization, error) {
	return s.repository.FindByName(ctx, name)
}
