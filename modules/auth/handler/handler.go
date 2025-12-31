package handler

import (
	"urlshortener/internal/database"
	"urlshortener/internal/models"
	"urlshortener/internal/utils"
	"urlshortener/modules/auth/dto"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var input dto.RegisterDTO
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"error": "Can not parse JSON.",
			},
		)
	}
	if errors := utils.ValidateStruct(input); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"error": errors,
			},
		)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), 12)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"error": "Hashing of password failded.",
			},
		)
	}
	user := models.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  string(hash),
	}

	if result := database.DB.Create(&user); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"error": "Failed to create the user",
			},
		)
	}

	return c.JSON(fiber.Map{
		"message": "User created.",
	})
}

func Login(c *fiber.Ctx) error {

	var input dto.LoginDTO
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"error": "Can not parse JSON.",
			},
		)
	}
	if errors := utils.ValidateStruct(input); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"error": errors,
			},
		)
	}

	var user models.User

	database.DB.Where("email = ?", input.Email).First(&user)

	if user.ID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"error": "Invalid email or password.",
			},
		)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"error": "Invalid email or password.",
			},
		)
	}

	secret := viper.GetString("JWT_SECRET")
	token, err := utils.GenerateJWTToken(user.ID, secret)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"error": "Failed to login",
			},
		)
	}

	return c.JSON(fiber.Map{
		"user_id": user.ID,
		"token":   token,
	})
}
