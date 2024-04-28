package repository

import (
	"context"
	"fmt"

	"nemesis-cli/src/models"

	"gorm.io/gorm"
)

type OrganizationRepository interface {
	GetAll(ctx context.Context) ([]models.Organization, error)
	GetByID(ctx context.Context, id string) (models.Organization, error)
	FindByName(ctx context.Context, name string) ([]models.Organization, error)
}

type organizationRepository struct {
	db *gorm.DB
}

func NewOrganizationRepository(db *gorm.DB) OrganizationRepository {
	return &organizationRepository{
		db: db,
	}
}

func (r *organizationRepository) GetAll(ctx context.Context) ([]models.Organization, error) {
	var organizations []models.Organization
	err := r.db.WithContext(ctx).Find(&organizations).Error
	if err != nil {
		return nil, err
	}
	return organizations, nil
}

func (r *organizationRepository) GetByID(ctx context.Context, id string) (models.Organization, error) {
	var organization models.Organization
	err := r.db.WithContext(ctx).First(&organization, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.Organization{}, err
		}
		return models.Organization{}, err
	}
	return organization, nil
}

func (r *organizationRepository) FindByName(ctx context.Context, name string) ([]models.Organization, error) {
	var organization []models.Organization
	err := r.db.WithContext(ctx).Where("name LIKE ?", fmt.Sprintf("%s%s%s", "%", name, "%")).Find(&organization).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return []models.Organization{}, err
		}
		return []models.Organization{}, err
	}
	return organization, nil
}
