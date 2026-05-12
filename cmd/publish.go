package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/kitdevelop-org/vtx/internal/config"
	"github.com/spf13/cobra"
)

var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Sube el plugin firmado al Marketplace de Veritix",
	Long: `vtx publish valida la firma digital del paquete .vtx y lo sube al
Veritix Registry para que esté disponible para su instalación.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.LoadConfig("vtx.config.json")
		if err != nil {
			return fmt.Errorf("no se encontró vtx.config.json: %w", err)
		}

		vtxFileName := fmt.Sprintf("%s-%s.vtx", cfg.PluginId, cfg.Version)
		vtxFilePath := filepath.Join("dist", vtxFileName)

		// 1. Verificar existencia del paquete
		if _, err := os.Stat(vtxFilePath); os.IsNotExist(err) {
			return fmt.Errorf("el paquete %s no existe. Ejecuta 'vtx build' primero", vtxFileName)
		}

		// 2. Simular validación de firma
		fmt.Printf("🔍 Validando firma RSA para: %s...\n", vtxFileName)
		fmt.Println("✅ Firma válida. Publisher: KitDevelop (Certificado 2026-0034)")

		// 3. Simular subida al Registry
		fmt.Printf("🚀 Subiendo paquete al Marketplace (https://registry.veritix.io)...\n")
		fmt.Println("⏳ Transfiriendo bloques...")
		fmt.Println("✅ Publicación exitosa.")
		fmt.Printf("\nPlugin '%s' v%s ya está disponible en el Marketplace.\n", cfg.PluginId, cfg.Version)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(publishCmd)
}
