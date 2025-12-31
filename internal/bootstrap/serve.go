package bootstrap

import (
	"log"
	"urlshortener/internal/config"

	"github.com/gofiber/fiber/v2"
)

func StartServer() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!!!!!!!!!!")
	})

	log.Fatal(app.Listen(":" + cfg.Port))
}
