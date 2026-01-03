package routes

import (
	"urlshortener/internal/middleware"
	"urlshortener/modules/link/handler"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoute(router fiber.Router) {

	linkGroup := router.Group("/link")

	linkGroup.Post("/", middleware.AuthMiddleware, handler.CreateLink)
	linkGroup.Delete("/:id", middleware.AuthMiddleware, handler.DeleteLink)
	linkGroup.Get("/", middleware.AuthMiddleware, handler.GetAllLinks)

	router.Get("/:code", handler.RedirectToLink)
}
