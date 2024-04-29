package repository

import (
	"context"
	"fmt"

	"gorm.io/gorm"
	"nemesis-cli/src/models"
)

type NetworkRepository interface {
	GetAll(ctx context.Context) ([]models.Networks, error)
	GetByID(ctx context.Context, id string) (models.Networks, error)
	FindByName(ctx context.Context, name string) ([]models.Networks, error)
}

type networkRepository struct {
	db *gorm.DB
}

func NewNetworkRepository(db *gorm.DB) NetworkRepository {
	return &networkRepository{
		db: db,
	}
}

func (r *networkRepository) GetAll(ctx context.Context) ([]models.Networks, error) {
	var networks []models.Networks

	err := r.db.WithContext(ctx).Find(&networks).Error
	if err != nil {
		return nil, err
	}
	return networks, nil
}

func (r *networkRepository) GetByID(ctx context.Context, id string) (models.Networks, error) {
	var network models.Networks
	err := r.db.WithContext(ctx).First(&network, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.Networks{}, err
		}
		return models.Networks{}, err
	}
	return network, nil
}

func (r *networkRepository) FindByName(ctx context.Context, name string) ([]models.Networks, error) {
	var network []models.Networks
	err := r.db.WithContext(ctx).Where("name LIKE ?", fmt.Sprintf("%s%s%s", "%", name, "%")).Find(&network).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return []models.Networks{}, err
		}
		return []models.Networks{}, err
	}
	return network, nil
}
