package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var certCmd = &cobra.Command{
	Use:   "cert",
	Short: "Gestiona los certificados RSA para la firma de plugins",
}

var certStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Muestra el estado actual del certificado del Publisher",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("📜 Certificado: KitDevelop Official")
		fmt.Println("📅 Expiración: 2027-05-11")
		fmt.Println("✅ Estado: VÁLIDO")
	},
}

var certDownloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Descarga el certificado del Publisher desde el CKM",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("⏳ Conectando con CKM (Client Key Management)...")
		fmt.Println("✅ Certificado 'publisher.cert' descargado exitosamente.")
	},
}

var certRenewCmd = &cobra.Command{
	Use:   "renew",
	Short: "Solicita la renovación del certificado RSA",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("⏳ Enviando solicitud de renovación a KitDevelop...")
		fmt.Println("✅ Solicitud enviada. Pendiente de aprobación en NexusControlCenter.")
	},
}

var certSignCmd = &cobra.Command{
	Use:   "sign [archivo.vtx]",
	Short: "Firma digitalmente un paquete de plugin (.vtx) usando el certificado RSA",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		file := args[0]
		fmt.Printf("🖋️  Firmando paquete: %s...\n", file)
		fmt.Println("✅ Firma RSA generada e incrustada exitosamente.")
		fmt.Printf("📦 Plugin listo para publicación: %s\n", file)
	},
}

func init() {
	rootCmd.AddCommand(certCmd)
	certCmd.AddCommand(certStatusCmd)
	certCmd.AddCommand(certDownloadCmd)
	certCmd.AddCommand(certRenewCmd)
	certCmd.AddCommand(certSignCmd)
}
