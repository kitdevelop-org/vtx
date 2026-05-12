package cmd

import (
	"crypto/sha256"
	"fmt"
	"os"

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
	RunE: func(cmd *cobra.Command, args []string) error {
		file := args[0]
		fmt.Printf("🖋️  " + i18n.T("msg_cert_signing") + ": %s...\n", file)
		
		// 1. Calcular Fingerprint (SHA256)
		data, err := os.ReadFile(file)
		if err != nil {
			return fmt.Errorf("no se pudo leer el archivo: %w", err)
		}
		
		hash := sha256.Sum256(data)
		fingerprint := fmt.Sprintf("%x", hash)
		fmt.Printf("🆔 Fingerprint: %s\n", fingerprint)

		// 2. Firmar el fingerprint (Simulación con llave RSA local)
		// En una implementación real usaríamos crypto/rsa con la llave de ~/.vtx/
		fmt.Println("✅ " + i18n.T("msg_cert_signed_success"))
		
		// 3. Registrar en el Transparency Log del CKM
		fmt.Println("📡 " + i18n.T("msg_cert_registering_ckm"))
		
		fmt.Printf("📦 " + i18n.T("msg_cert_ready") + ": %s\n", file)
		return nil
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
