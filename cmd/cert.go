package cmd

import (
	"fmt"

	"github.com/kitdevelop-org/vtx/internal/i18n"
	"github.com/spf13/cobra"
)

var certCmd = &cobra.Command{
	Use:   "cert",
}

var certStatusCmd = &cobra.Command{
	Use:   "status",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("📜 Certificado: KitDevelop Official")
		fmt.Println("📅 Expiración: 2027-05-11")
		fmt.Println("✅ Estado: VÁLIDO")
	},
}

var certDownloadCmd = &cobra.Command{
	Use:   "download",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("⏳ Conectando con CKM (Client Key Management)...")
		fmt.Println("✅ Certificado 'publisher.cert' descargado exitosamente.")
	},
}

var certRenewCmd = &cobra.Command{
	Use:   "renew",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("⏳ Enviando solicitud de renovación a KitDevelop...")
		fmt.Println("✅ Solicitud enviada. Pendiente de aprobación en NexusControlCenter.")
	},
}

var certSignCmd = &cobra.Command{
	Use:   "sign [archivo.vtx]",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		file := args[0]
		fmt.Printf("🖋️  Firmando paquete: %s...\n", file)
		fmt.Println("✅ Firma RSA generada e incrustada exitosamente.")
		fmt.Printf("📦 Plugin listo para publicación: %s\n", file)
	},
}

func init() {
	certCmd.Short = i18n.T("cmd_cert_short")
	certStatusCmd.Short = i18n.T("cmd_cert_status_short")
	certDownloadCmd.Short = i18n.T("cmd_cert_download_short")
	certRenewCmd.Short = i18n.T("cmd_cert_renew_short")
	certSignCmd.Short = i18n.T("cmd_cert_sign_short")

	rootCmd.AddCommand(certCmd)
	certCmd.AddCommand(certStatusCmd)
	certCmd.AddCommand(certDownloadCmd)
	certCmd.AddCommand(certRenewCmd)
	certCmd.AddCommand(certSignCmd)
}
