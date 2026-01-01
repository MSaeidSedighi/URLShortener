package middleware

import (
	"fmt"
	"strings"
	"urlshortener/internal/constants"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

func AuthMiddleware(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization", "")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(
			fiber.Map{
				"error": "Unathorized Access.",
			},
		)
	}

	jwtSecret := viper.GetString("JWT_SECRET")

	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(jwtSecret), nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(
			fiber.Map{
				"error": "Invalid token",
			},
		)
	}

	userId := cast.ToUint(claims["user_id"])

	c.Locals(constants.USER_ID_KEY, userId)

	return c.Next()
}
