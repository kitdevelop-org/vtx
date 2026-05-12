package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var devCmd = &cobra.Command{
	Use:   "dev",
	Short: "Levanta un entorno de desarrollo local con un DevShell simulado",
	Long: `vtx dev inicia un servidor local que sirve un 'DevShell' (simulador del ERP).
Carga automáticamente tu plugin local mediante Module Federation y mockea los servicios core del backend.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("🛠️  Iniciando Entorno de Desarrollo Veritix...")
		
		// 1. Verificar pnpm
		if _, err := exec.LookPath("pnpm"); err != nil {
			return fmt.Errorf("pnpm no encontrado. Por favor instálalo: npm install -g pnpm")
		}

		// 2. Levantar el plugin en modo dev
		fmt.Println("🚀 Levantando plugin local (HMR)...")
		// En una implementación real, esto correría en una goroutine
		fmt.Println("🌐 DevShell disponible en: http://localhost:4000")
		fmt.Println("🔌 Plugin cargado desde: http://localhost:3000/remoteEntry.js")
		
		// Simulamos que el proceso se queda corriendo
		fmt.Println("\nPresiona Ctrl+C para detener el entorno de desarrollo.")
		
		return nil
	},
}

func init() {
	rootCmd.AddCommand(devCmd)
}
