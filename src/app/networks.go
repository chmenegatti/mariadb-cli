package app

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/aquasecurity/table"
	"github.com/liamg/tml"
	"nemesis-cli/src/handler"
	"nemesis-cli/src/models"
	"nemesis-cli/src/repository"
	"nemesis-cli/src/service"
	"nemesis-cli/src/utils"
)

func Networks(id, name string) error {

	netRepo := repository.NewNetworkRepository(dbInstance)
	netService := service.NewNetworkService(netRepo)

	t.SetDividers(table.UnicodeRoundedDividers)
	t.SetLineStyle(table.StyleBlue)

	headers := utils.ParseHeader(models.Networks{})

	t.SetHeaders(headers...)

	if id == "" && name == "" {
		if err := allNets(netService); err != nil {
			return fmt.Errorf("error: %w", err)
		}
	}

	if id != "" {
		if err := byIdNets(netService, id); err != nil {
			return fmt.Errorf("error: %w", err)
		}
	}

	if name != "" {
		if err := byNameNets(netService, name); err != nil {
			return fmt.Errorf("error: %w", err)
		}
	}

	t.Render()

	return nil
}

func allNets(netService service.NetworkService) error {
	nets, err := handler.GetAllNetworks(context.Background(), netService)
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}

	for _, net := range nets {
		parseFieldsNetworks(net)
	}

	return nil
}

func byIdNets(netService service.NetworkService, id string) error {
	net, err := handler.GetByIdNetworks(context.Background(), netService, id)
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}

	parseFieldsNetworks(net)

	return nil
}

func byNameNets(netService service.NetworkService, name string) error {
	nets, err := handler.GetByNameNetworks(context.Background(), netService, name)
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}
	for _, net := range nets {

		if strings.Contains(strings.ToLower(net.Name), strings.ToLower(name)) {
			net.Name = strings.ReplaceAll(net.Name, name, tml.Sprintf("<green>%s</green>", name))
		}
		parseFieldsNetworks(net)
	}

	return nil
}

func parseFieldsNetworks(net models.Networks) {

	errorString := utils.ParseError(net.Error)
	parsedDate := utils.ParseDate(net.Created.String())

	t.AddRow(
		net.Id, net.Name, net.Description, net.Address, strconv.FormatBool(net.EnableSideCommunication), net.Status,
		errorString,
		parsedDate,
	)

}
