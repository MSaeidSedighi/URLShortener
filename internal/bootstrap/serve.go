package bootstrap

import (
	"log"
	"urlshortener/internal/config"
	"urlshortener/internal/database"
	authRouter "urlshortener/modules/auth/routes"

	"github.com/gofiber/fiber/v2"
)

func StartServer() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	database.Connect(cfg.DbUrl)

	app := fiber.New()

	authRouter.RegisterRoute(app)

	log.Fatal(app.Listen(":" + cfg.Port))
}
