package app

import (
	"context"
	"os"
	"reflect"
	"time"

	"github.com/liamg/tml"
	"nemesis-cli/src/database"
	"nemesis-cli/src/handler"
	"nemesis-cli/src/models"
	"nemesis-cli/src/repository"
	"nemesis-cli/src/service"

	"github.com/aquasecurity/table"
)

func Organizations(id, name string) error {

	dbInstance := database.GetDB()
	orgsRepo := repository.NewOrganizationRepository(dbInstance)
	orgsService := service.NewOrganizationService(orgsRepo)

	var (
		t = table.New(os.Stdout)
	)

	t.SetDividers(table.UnicodeRoundedDividers)
	t.SetLineStyle(table.StyleBlue)
	x := reflect.TypeOf(models.Organization{})

	var headers []string

	for i := 0; i < x.NumField(); i++ {
		field := x.Field(i)
		if field.Name == "LoadBalanceSize" {
			field.Name = "LB Size"
		}
		if field.Name == "PhysicalFirewall" {
			field.Name = "Phys FW"
		}
		headers = append(headers, field.Name)
	}

	t.SetHeaders(headers...)

	if id == "" && name == "" {
		resp, err := handler.Index(context.Background(), orgsService)
		if err != nil {
			panic(err)
		}
		for _, org := range resp {
			var err string
			if org.Error != nil {
				err = *org.Error
				if len(err) > 15 {
					err = err[:15] + "..."
				}
			} else {
				err = ""
			}
			t.AddRow(
				org.ID, org.Name, org.TierProvider, org.BackupCluster, org.PhysicalFirewall,
				org.VirtualFirewall, org.LoadBalanceSize, org.Status, err, org.Created.String(),
			)

		}

		t.Render()

	}

	if id != "" {
		resp, err := handler.Show(context.Background(), orgsService, id)
		if err != nil {
			panic(err)
		}

		var erro string
		if resp.Error != nil {
			erro = *resp.Error
			if len(erro) > 15 {
				erro = erro[:15] + "..."
			}
		} else {
			erro = ""
		}

		t.AddRow(
			resp.ID, resp.Name, resp.TierProvider, resp.BackupCluster, resp.PhysicalFirewall,
			resp.VirtualFirewall, resp.LoadBalanceSize, resp.Status, erro, resp.Created.String(),
		)
		t.Render()
	}

	if name != "" {
		resp, err := handler.FindByName(context.Background(), orgsService, name)
		if err != nil {
			panic(err)
		}
		for _, org := range resp {
			var err string
			if org.Error != nil {
				err = *org.Error
				if len(err) > 15 {
					err = err[:15] + "..."
				}
			} else {
				err = ""
			}

			h, _ := time.Parse("2006-01-02 15:04:05 -0700 MST", org.Created.String())

			t.AddRow(
				org.ID, tml.Sprintf("<green>%s</green>", org.Name), org.TierProvider, org.BackupCluster, org.PhysicalFirewall,
				org.VirtualFirewall, org.LoadBalanceSize, org.Status, err, h.Format("02/01/06 15:04"),
			)

		}

		t.Render()
	}

	return nil
}
