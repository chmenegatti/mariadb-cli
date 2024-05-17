package repository

import (
	"context"
	"fmt"

	"gorm.io/gorm"
	"nemesis-cli/src/models"
)

type VirtualMachinesRepository interface {
	GetAll(ctx context.Context) ([]models.VirtualMachines, error)
	GetByID(ctx context.Context, id string) (models.VirtualMachines, error)
	FindByName(ctx context.Context, name string) ([]models.VirtualMachines, error)
	FindByOrganization(ctx context.Context, organization string) ([]models.VirtualMachines, error)
	FindByTopology(ctx context.Context, topology string) ([]models.VirtualMachines, error)
}

type virtualMachinesRepository struct {
	db *gorm.DB
}

func NewVirtualMachinesRepository(db *gorm.DB) VirtualMachinesRepository {
	return &virtualMachinesRepository{
		db: db,
	}
}

func (r *virtualMachinesRepository) GetAll(ctx context.Context) ([]models.VirtualMachines, error) {
	var virtualMachines []models.VirtualMachines
	err := r.db.WithContext(ctx).Find(&virtualMachines).Error
	if err != nil {
		return nil, err
	}
	return virtualMachines, nil
}

func (r *virtualMachinesRepository) GetByID(ctx context.Context, id string) (models.VirtualMachines, error) {
	var virtualMachines models.VirtualMachines
	err := r.db.WithContext(ctx).First(&virtualMachines, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.VirtualMachines{}, err
		}
		return models.VirtualMachines{}, err
	}
	return virtualMachines, nil
}

func (r *virtualMachinesRepository) FindByName(ctx context.Context, name string) ([]models.VirtualMachines, error) {
	var virtualMachines []models.VirtualMachines
	err := r.db.WithContext(ctx).Where("name LIKE ?", fmt.Sprintf("%s%s%s", "%", name, "%")).Find(&virtualMachines).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return []models.VirtualMachines{}, err
		}
		return []models.VirtualMachines{}, err
	}
	return virtualMachines, nil
}

func (r *virtualMachinesRepository) FindByOrganization(
	ctx context.Context, organizationID string,
) ([]models.VirtualMachines, error) {
	var virtualMachines []models.VirtualMachines
	err := r.db.WithContext(ctx).Where("organization = ?", organizationID).Find(&virtualMachines).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return []models.VirtualMachines{}, err
		}
		return []models.VirtualMachines{}, err
	}
	return virtualMachines, nil
}

func (r *virtualMachinesRepository) FindByTopology(
	ctx context.Context, topologyID string,
) ([]models.VirtualMachines, error) {
	var virtualMachines []models.VirtualMachines
	err := r.db.WithContext(ctx).Where("topology = ?", topologyID).Find(&virtualMachines).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return []models.VirtualMachines{}, err
		}
		return []models.VirtualMachines{}, err
	}
	return virtualMachines, nil
}
