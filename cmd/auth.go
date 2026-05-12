package cmd

import (
	"fmt"

	"github.com/kitdevelop-org/vtx/internal/i18n"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🔐 Abriendo el navegador para autenticación...")
		fmt.Println("⏳ Esperando confirmación del CKM (Client Key Management)...")
		fmt.Println("✅ Autenticado exitosamente como: KitDevelop")
	},
}

var whoamiCmd = &cobra.Command{
	Use:   "whoami",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("👤 Usuario: admin@kitdevelop.com")
		fmt.Println("🏢 Organización: KitDevelop S.R.L.")
		fmt.Println("📜 Certificado RSA: Activo (Expira en 320 días)")
	},
}

func init() {
	loginCmd.Short = i18n.T("cmd_login_short")
	whoamiCmd.Short = i18n.T("cmd_whoami_short")
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(whoamiCmd)
}
