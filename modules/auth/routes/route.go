package routes

import (
	"urlshortener/modules/auth/handler"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoute(router fiber.Router) {
	authenticationGroup := router.Group("/auth")

	authenticationGroup.Post("/register", handler.Register)
	authenticationGroup.Post("/login", handler.Login)

}
