package app

import (
	"context"
	"nemesis-cli/src/database"
	"nemesis-cli/src/handler"
	"nemesis-cli/src/repository"
	"nemesis-cli/src/service"
	"os"

	"github.com/aquasecurity/table"
	"github.com/liamg/tml"
)

func Organizations(id, name string) error {

	dbInstance := database.GetDB()
	orgsRepo := repository.NewOrganizationRepository(dbInstance)
	orgsService := service.NewOrganizationService(orgsRepo)

	var (
		t = table.New(os.Stdout)
	)

	t.SetHeaders("ID", "Name", "Description")

	if id == "" && name == "" {
		resp, err := handler.Index(context.Background(), orgsService)
		if err != nil {
			panic(err)
		}

		for _, org := range resp {
			t.AddRow(org.ID, org.Name, org.Description)
		}

		t.Render()

	}

	if id != "" {
		resp, err := handler.Show(context.Background(), orgsService, id)
		if err != nil {
			panic(err)
		}

		t.AddRow(resp.ID, resp.Name, resp.Description)
		t.Render()
	}

	if name != "" {
		resp, err := handler.FindByName(context.Background(), orgsService, name)
		if err != nil {
			panic(err)
		}

		for _, org := range resp {
			t.AddRow(org.ID, tml.Sprintf("<green>%s</green>", org.Name), org.Description)
		}

		t.Render()
	}

	return nil
}
