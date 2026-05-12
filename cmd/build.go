package cmd

import (
	"fmt"

	"github.com/kitdevelop-org/vtx/internal/builder"
	"github.com/kitdevelop-org/vtx/internal/config"
	"github.com/kitdevelop-org/vtx/internal/i18n"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: i18n.T("cmd_build_short"),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.LoadConfig("vtx.config.json")
		if err != nil {
			return fmt.Errorf("no se encontró vtx.config.json: %w", err)
		}

		if err := builder.Build(cfg); err != nil {
			return err
		}

		fmt.Println("\n✅ Build completado exitosamente.")
		fmt.Println("🔑 Pendiente: Firma RSA mediante 'vtx cert sign' (Fase CKM)")
		return nil
	},
}

func init() {
	buildCmd.Short = i18n.T("cmd_build_short")
	rootCmd.AddCommand(buildCmd)
}
