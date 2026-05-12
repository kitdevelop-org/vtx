package ui

import (
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/kitdevelop-org/vtx/internal/i18n"
)

type NewPluginAnswers struct {
	Name          string
	PluginId      string
	HasFrontend   bool
	HasContracts  bool
	Country       string
	Author        string
	LicenseType   string // "Free", "Veritix", "Custom"
}

func RunNewWizard(defaultName string) (*NewPluginAnswers, error) {
	answers := &NewPluginAnswers{
		Name: defaultName,
	}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title(i18n.T("wizard_name_title")).
				Placeholder("Ej: Inventory, POS Retail").
				Value(&answers.Name).
				Validate(func(s string) error {
					if s == "" {
						return fmt.Errorf("el nombre es obligatorio")
					}
					return nil
				}),

			huh.NewInput().
				Title(i18n.T("wizard_id_title")).
				Placeholder("Ej: MyInventory").
				Value(&answers.PluginId),

			huh.NewConfirm().
				Title(i18n.T("wizard_frontend_title")).
				Value(&answers.HasFrontend),

			huh.NewConfirm().
				Title(i18n.T("wizard_contracts_title")).
				Value(&answers.HasContracts),

			huh.NewSelect[string]().
				Title(i18n.T("wizard_license_title")).
				Options(
					huh.NewOption(i18n.T("license_free"), "Free"),
					huh.NewOption(i18n.T("license_veritix"), "Veritix"),
					huh.NewOption(i18n.T("license_custom"), "Custom"),
				).
				Value(&answers.LicenseType),

			huh.NewSelect[string]().
				Title(i18n.T("wizard_country_title")).
				Options(
					huh.NewOption("República Dominicana (DO)", "DO"),
					huh.NewOption("Colombia (CO)", "CO"),
					huh.NewOption("México (MX)", "MX"),
					huh.NewOption("Global / Ninguno", "GLOBAL"),
				).
				Value(&answers.Country),

			huh.NewInput().
				Title(i18n.T("wizard_author_title")).
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
