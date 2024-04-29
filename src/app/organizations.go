package app

import (
	"context"
	"fmt"
	"strings"

	"github.com/liamg/tml"
	"nemesis-cli/src/handler"
	"nemesis-cli/src/models"
	"nemesis-cli/src/repository"
	"nemesis-cli/src/service"
	"nemesis-cli/src/utils"

	"github.com/aquasecurity/table"
)

func Organizations(id, name string) error {

	orgsRepo := repository.NewOrganizationRepository(dbInstance)
	orgsService := service.NewOrganizationService(orgsRepo)

	t.SetDividers(table.UnicodeRoundedDividers)
	t.SetLineStyle(table.StyleBlue)

	headers := utils.ParseHeader(models.Organization{})

	t.SetHeaders(headers...)

	if id == "" && name == "" {
		if err := allOrgs(orgsService); err != nil {
			return fmt.Errorf("error: %w", err)
		}
	}

	if id != "" {
		if err := byIdOrgs(orgsService, id); err != nil {
			return fmt.Errorf("error: %w", err)
		}
	}

	if name != "" {
		if err := byNameOrgs(orgsService, name); err != nil {
			return fmt.Errorf("error: %w", err)
		}
	}

	t.Render()
	return nil
}

func allOrgs(orgService service.OrganizationService) error {
	resp, err := handler.GetAllOrganizations(context.Background(), orgService)
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}
	for _, org := range resp {
		errorString := utils.ParseError(org.Error)
		parsedDate := utils.ParseDate(org.Created.String())
		t.AddRow(
			org.ID, org.Name, org.TierProvider, org.BackupCluster, org.PhysicalFirewall,
			org.VirtualFirewall, org.LoadBalanceSize, org.Status, errorString, parsedDate,
		)
	}
	return nil
}

func byIdOrgs(orgService service.OrganizationService, id string) error {
	org, err := handler.GetByIdOrganization(context.Background(), orgService, id)
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}

	errorString := utils.ParseError(org.Error)
	parsedDate := utils.ParseDate(org.Created.String())

	t.AddRow(
		org.ID, org.Name, org.TierProvider, org.BackupCluster, org.PhysicalFirewall,
		org.VirtualFirewall, org.LoadBalanceSize, org.Status, errorString, parsedDate,
	)

	return nil
}

func byNameOrgs(orgService service.OrganizationService, name string) error {
	resp, err := handler.GetByNameOrganization(context.Background(), orgService, name)
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}
	for _, org := range resp {
		errorString := utils.ParseError(org.Error)

		if strings.Contains(strings.ToLower(org.Name), strings.ToLower(name)) {
			org.Name = strings.ReplaceAll(org.Name, name, tml.Sprintf("<green>%s</green>", name))
		}

		t.AddRow(
			org.ID, org.Name, org.TierProvider, org.BackupCluster, org.PhysicalFirewall,
			org.VirtualFirewall, org.LoadBalanceSize, org.Status, errorString, org.Created.String(),
		)
	}

	return nil
}
