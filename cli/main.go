package main

import (
	"os"

	"github.com/OLIENTTECH/backend-challenges/server"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "entrypoint",
		Short: "CLI tool for managing backend challenges server",
	}

	rootCmd.AddCommand(
		server.NewCommand(),
	)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
