package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vtx",
	Short: "vtx es la CLI oficial para desarrollar plugins de Veritix ERP",
	Long: `vtx automatiza el ciclo de vida completo del desarrollo de plugins para Veritix.
Desde el scaffolding inicial hasta la publicación firmada en el Marketplace.`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Aquí se pueden añadir flags globales
}
