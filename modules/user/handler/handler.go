package handler

import (
	"urlshortener/internal/constants"
	"urlshortener/internal/database"
	"urlshortener/internal/models"
	"urlshortener/internal/utils"
	"urlshortener/modules/user/dto"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cast"
)

func GetProfile(c *fiber.Ctx) error {
	userId := cast.ToUint(c.Locals(constants.USER_ID_KEY))

	var user models.User
	database.DB.Where("id = ?", userId).First(&user)

	if userId == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"error": "User not found.",
			},
		)
	}
	return c.JSON(fiber.Map{
		"id":         user.ID,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"email":      user.Email,
		"created_at": user.CreatedAt,
		"updated_at": user.UpdatedAt,
	})
}

func UpdateProfile(c *fiber.Ctx) error {
	var input dto.UpdateUserDTO
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

	userId := cast.ToUint(c.Locals(constants.USER_ID_KEY))

	var user models.User
	database.DB.Where("id = ?", userId).First(&user)

	if userId == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"error": "User not found.",
			},
		)
	}

	user.FirstName = input.FirstName
	user.LastName = input.LastName

	if err := database.DB.Save(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"error": "Failed to update user",
			},
		)
	}

	return c.JSON(fiber.Map{
		"id":         user.ID,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"email":      user.Email,
		"created_at": user.CreatedAt,
		"updated_at": user.UpdatedAt,
	})
}
