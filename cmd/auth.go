package cmd

import (
	"fmt"

	"github.com/kitdevelop-org/vtx/internal/i18n"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("🔐 " + i18n.T("msg_auth_opening_browser"))
		
		// Simulación de llamada al CKM
		// En producción esto usaría el endpoint /publishers/auth/keys
		fmt.Println("⏳ " + i18n.T("msg_auth_waiting_ckm"))
		
		// Guardamos las llaves simuladas localmente en ~/.vtx/
		home, _ := os.UserHomeDir()
		vtxDir := filepath.Join(home, ".vtx")
		os.MkdirAll(vtxDir, 0700)

		fmt.Println("✅ " + i18n.T("msg_auth_success") + ": KitDevelop")
		fmt.Println("🔑 " + i18n.T("msg_auth_keys_downloaded"))
		
		return nil
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
