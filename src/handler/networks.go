package handler

import (
	"context"

	"nemesis-cli/src/models"
	"nemesis-cli/src/service"
)

func GetAllNetworks(ctx context.Context, service service.NetworkService) (
	nets []models.Networks, err error,
) {
	var networks []models.Networks

	networks, err = service.GetAll(ctx)
	if err != nil {
		return
	}

	return networks, nil
}

func GetByIdNetworks(ctx context.Context, service service.NetworkService, id string) (
	net models.Networks, err error,
) {
	var network models.Networks

	network, err = service.GetByID(ctx, id)

	if err != nil {
		return
	}

	return network, nil
}

func GetByNameNetworks(
	ctx context.Context, service service.NetworkService, name string,
) (nets []models.Networks, err error) {
	var networks []models.Networks
	networks, err = service.GetByName(ctx, name)
	if err != nil {
		return
	}

	return networks, nil
}
