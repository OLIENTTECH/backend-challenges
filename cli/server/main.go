package main

import (
	"os"

	"github.com/OLIENTTECH/backend-challenges/gateway"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "server",
		Short: "CLI tool for managing DX Henpin services",
	}

	rootCmd.AddCommand(
		gateway.NewCommand(),
	)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
