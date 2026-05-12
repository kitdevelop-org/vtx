package ui

import (
	"fmt"

	"github.com/charmbracelet/huh"
)

type NewPluginAnswers struct {
	Name          string
	PluginId      string
	HasFrontend   bool
	HasContracts  bool
	Country       string
	Author        string
}

func RunNewWizard(defaultName string) (*NewPluginAnswers, error) {
	answers := &NewPluginAnswers{
		Name: defaultName,
	}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("¿Cuál es el nombre del plugin?").
				Placeholder("Ej: Inventory, POS Retail").
				Value(&answers.Name).
				Validate(func(s string) error {
					if s == "" {
						return fmt.Errorf("el nombre es obligatorio")
					}
					return nil
				}),

			huh.NewInput().
				Title("Plugin ID (PascalCase)").
				Placeholder("Ej: MyInventory").
				Value(&answers.PluginId),

			huh.NewConfirm().
				Title("¿Incluir proyecto Frontend (ClientApp)?").
				Value(&answers.HasFrontend),

			huh.NewConfirm().
				Title("¿Deseas que este plugin sea utilizable como dependencia? (Crea proyecto .Contracts)").
				Value(&answers.HasContracts),

			huh.NewSelect[string]().
				Title("País Fiscal base").
				Options(
					huh.NewOption("República Dominicana (DO)", "DO"),
					huh.NewOption("Colombia (CO)", "CO"),
					huh.NewOption("México (MX)", "MX"),
					huh.NewOption("Global / Ninguno", "GLOBAL"),
				).
				Value(&answers.Country),

			huh.NewInput().
				Title("Autor / Organización").
				Placeholder("Ej: KitDevelop").
				Value(&answers.Author),
		),
	)

	err := form.Run()
	if err != nil {
		return nil, err
	}

	return answers, nil
}
