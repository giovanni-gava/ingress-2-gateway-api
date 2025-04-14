package cli

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// NewRootCommand monta o comando raiz da CLI
func NewRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "gateway-migrator",
		Short: "Migrate safely from Ingress to Gateway API",
		Long:  `CLI tool to automate and validate Kubernetes Ingress to Gateway API migration`,
	}

	cobra.OnInitialize(initConfig)

	// Registra os subcomandos (como 'diff')
	rootCmd.AddCommand(NewDiffCommand())

	return rootCmd
}

func initConfig() {
	viper.AutomaticEnv()
}
