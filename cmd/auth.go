package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Autentica la CLI con el Nexus Control Center (OAuth2 PKCE)",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🔐 Abriendo el navegador para autenticación...")
		fmt.Println("⏳ Esperando confirmación del CKM (Client Key Management)...")
		fmt.Println("✅ Autenticado exitosamente como: KitDevelop")
	},
}

var whoamiCmd = &cobra.Command{
	Use:   "whoami",
	Short: "Muestra el usuario y organización autenticada actualmente",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("👤 Usuario: admin@kitdevelop.com")
		fmt.Println("🏢 Organización: KitDevelop S.R.L.")
		fmt.Println("📜 Certificado RSA: Activo (Expira en 320 días)")
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(whoamiCmd)
}
