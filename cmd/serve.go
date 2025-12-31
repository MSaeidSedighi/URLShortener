package cmd

import (
	"urlshortener/internal/bootstrap"

	"github.com/spf13/cobra"
)

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the HTTP Server",
	Run: func(cmd *cobra.Command, args []string) {
		bootstrap.StartServer()
	},
}
