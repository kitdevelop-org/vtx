package cmd

import (
	"github.com/kitdevelop-org/vtx/internal/i18n"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "vtx",
}

func Execute() error {
	// 1. Configurar descripciones del comando raíz
	rootCmd.Short = i18n.T("cli_short_desc")
	rootCmd.Long = i18n.T("cli_long_desc")

	// 2. Configurar descripciones de subcomandos
	newCmd.Short = i18n.T("cmd_new_short")
	buildCmd.Short = i18n.T("cmd_build_short")
	publishCmd.Short = i18n.T("cmd_publish_short")
	devCmd.Short = i18n.T("cmd_dev_short")
	loginCmd.Short = i18n.T("cmd_login_short")
	whoamiCmd.Short = i18n.T("cmd_whoami_short")
	certCmd.Short = i18n.T("cmd_cert_short")
	certStatusCmd.Short = i18n.T("cmd_cert_status_short")
	certDownloadCmd.Short = i18n.T("cmd_cert_download_short")
	certRenewCmd.Short = i18n.T("cmd_cert_renew_short")
	certSignCmd.Short = i18n.T("cmd_cert_sign_short")
	addCmd.Short = i18n.T("cmd_add_short")

	return rootCmd.Execute()
}

func init() {
	// El registro ya ocurre en los archivos individuales mediante init()
}
