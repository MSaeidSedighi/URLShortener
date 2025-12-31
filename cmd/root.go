package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func Execute() {
	var rootCmd = &cobra.Command{
		Use:   "shortener",
		Short: "A URL Shortener CLI",
	}

	rootCmd.AddCommand(ServeCmd)

	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}

}
