package routes

import (
	"urlshortener/internal/middleware"
	"urlshortener/modules/user/handler"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoute(router fiber.Router) {
	userGroup := router.Group("/user")

	userGroup.Get("/profile", middleware.AuthMiddleware, handler.GetProfile)
	userGroup.Patch("/profile", middleware.AuthMiddleware, handler.UpdateProfile)

}
