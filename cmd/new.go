package cmd

import (
	"fmt"

	"github.com/kitdevelop-org/vtx/internal/scaffold"
	"github.com/kitdevelop-org/vtx/internal/ui"
	"github.com/spf13/cobra"
)

var (
	flagName     string
	flagId       string
	flagFrontend bool
	flagCountry  string
	flagAuthor   string
)

var newCmd = &cobra.Command{
	Use:   "new [nombre]",
	Short: "Crea un nuevo proyecto de plugin con la estructura oficial",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var answers *ui.NewPluginAnswers
		var err error

		// Si se pasan flags, saltar el wizard
		if flagName != "" && flagId != "" {
			answers = &ui.NewPluginAnswers{
				Name:        flagName,
				PluginId:    flagId,
				HasFrontend: flagFrontend,
				Country:     flagCountry,
				Author:      flagAuthor,
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

		fmt.Printf("\n🚀 Creando plugin: %s...\n", answers.Name)

		if err := scaffold.Generate(answers); err != nil {
			return fmt.Errorf("error generando el plugin: %w", err)
		}

		fmt.Println("✅ Estructura de carpetas creada exitosamente.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
	newCmd.Flags().StringVarP(&flagName, "name", "n", "", "Nombre del plugin")
	newCmd.Flags().StringVarP(&flagId, "id", "i", "", "ID del plugin (PascalCase)")
	newCmd.Flags().BoolVarP(&flagFrontend, "frontend", "f", true, "¿Incluir frontend?")
	newCmd.Flags().StringVarP(&flagCountry, "country", "c", "GLOBAL", "País fiscal")
	newCmd.Flags().StringVarP(&flagAuthor, "author", "a", "KitDevelop", "Autor/Organización")
}
