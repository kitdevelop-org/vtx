package cmd

import (
	"fmt"

	"github.com/kitdevelop-org/vtx/internal/i18n"
	"github.com/kitdevelop-org/vtx/internal/scaffold"
	"github.com/kitdevelop-org/vtx/internal/ui"
	"github.com/spf13/cobra"
)

var (
	flagName      string
	flagId        string
	flagFrontend  bool
	flagContracts bool
	flagCountry   string
	flagAuthor    string
	flagLicense   string
)

var newCmd = &cobra.Command{
	Use:   "new [nombre]",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var answers *ui.NewPluginAnswers
		var err error

		// Si se pasan flags, saltar el wizard
		if flagName != "" && flagId != "" {
			answers = &ui.NewPluginAnswers{
				Name:         flagName,
				PluginId:     flagId,
				HasFrontend:  flagFrontend,
				HasContracts: flagContracts,
				Country:      flagCountry,
				Author:       flagAuthor,
				LicenseType:  flagLicense,
			}
		} else {
			name := ""
			if len(args) > 0 {
				name = args[0]
			}
			answers, err = ui.RunNewWizard(name)
			if err != nil {
				return err
			}
		}

		fmt.Printf("\n"+i18n.T("msg_creating_plugin")+"\n", answers.Name)

		if err := scaffold.Generate(answers); err != nil {
			return fmt.Errorf("error generando el plugin: %w", err)
		}

		fmt.Println(i18n.T("msg_scaffold_success"))
		return nil
	},
}

func init() {
	newCmd.Short = i18n.T("cmd_new_short")
	rootCmd.AddCommand(newCmd)
	newCmd.Flags().StringVarP(&flagName, "name", "n", "", "Nombre del plugin")
	newCmd.Flags().StringVarP(&flagId, "id", "i", "", "ID del plugin (PascalCase)")
	newCmd.Flags().BoolVarP(&flagFrontend, "frontend", "f", true, "¿Incluir frontend?")
	newCmd.Flags().BoolVarP(&flagContracts, "contracts", "x", true, "¿Crear proyecto de contratos?")
	newCmd.Flags().StringVarP(&flagCountry, "country", "c", "GLOBAL", "País fiscal")
	newCmd.Flags().StringVarP(&flagAuthor, "author", "a", "KitDevelop", "Autor/Organización")
	newCmd.Flags().StringVarP(&flagLicense, "license", "l", "Free", "Tipo de licencia (Free, Veritix, Custom)")
}
