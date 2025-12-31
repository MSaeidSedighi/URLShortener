package cmd

import (
	"log"
	"urlshortener/internal/config"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
)

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the HTTP Server",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig()
		if err != nil {
			log.Fatal(err)
		}
		app := fiber.New()

		app.Get("/", func(c *fiber.Ctx) error {
			return c.SendString("Hello, World!!!!!!!!!!")
		})

		log.Fatal(app.Listen(":" + cfg.Port))
	},
}
