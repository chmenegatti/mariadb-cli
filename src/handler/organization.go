package handler

import (
	"context"
	"encoding/json"

	"nemesis-cli/src/models"
	"nemesis-cli/src/service"
)

func Index(ctx context.Context, service service.OrganizationService) (orgs []models.Organization, err error) {
	var organizations []models.Organization

	organizations, err = service.GetAll(ctx)
	if err != nil {
		return
	}

	return organizations, nil
}

func Show(ctx context.Context, service service.OrganizationService, id string) (string, error) {
	var organization models.Organization
	var err error

	organization, err = service.GetByID(ctx, id)
	if err != nil {
		return "", err
	}

	jsonData, err := json.Marshal(organization)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

func FindByName(ctx context.Context, service service.OrganizationService, name string) (orgs []models.Organization, err error) {
	var organizations []models.Organization

	organizations, err = service.GetByName(ctx, name)
	if err != nil {
		return
	}

	return organizations, nil
}
