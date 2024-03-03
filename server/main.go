package main

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/OLIENTTECH/backend-challenges/cli"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "entrypoint",
		Short: "CLI tool for managing backend challenges server",
	}

	rootCmd.AddCommand(
		cli.NewCommand(),
	)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
