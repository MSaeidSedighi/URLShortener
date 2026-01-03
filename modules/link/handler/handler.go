package handler

import (
	"urlshortener/internal/constants"
	"urlshortener/internal/database"
	"urlshortener/internal/models"
	"urlshortener/internal/utils"
	"urlshortener/modules/link/dto"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cast"
)

func RedirectToLink(c *fiber.Ctx) error {
	shortCode := c.Params("code")
	var link models.Link

	database.DB.Where("short_code = ?", shortCode).First(&link)
	if link.ID == 0 {
		return c.SendString("URL Not Found.")
	}

	database.DB.Model(&link).Update("visits", link.Visits+1)

	return c.Redirect(link.OriginalUrl, fiber.StatusTemporaryRedirect)
}

func CreateLink(c *fiber.Ctx) error {
	var input dto.CreateLinkDTO
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

	var shortCode string
	for {
		code := utils.GenerateRandomString(6)
		var codeCount int64
		database.DB.Model(&models.Link{}).Where("short_code = ?", code).Count(&codeCount)

		if codeCount == 0 {
			shortCode = code
			break
		}
	}

	link := models.Link{
		OriginalUrl: input.OriginalUrl,
		ShortCode:   shortCode,
		UserID:      userId,
		Visits:      0,
	}

	if result := database.DB.Create(&link); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"error": "Failed to create link.",
			},
		)
	}

	return c.JSON(fiber.Map{
		"short_code":   link.ShortCode,
		"original_url": input.OriginalUrl,
	})
}

func GetAllLinks(c *fiber.Ctx) error {
	userId := cast.ToUint(c.Locals(constants.USER_ID_KEY))

	var links []models.Link
	database.DB.Select("id, original_url, short_code, visits, created_at").
		Where("user_id = ?", userId).Find(&links)

	response := []dto.LinkResponse{}

	for _, link := range links {
		response = append(response, dto.LinkResponse{
			ID:          link.ID,
			OriginalUrl: link.OriginalUrl,
			ShortCode:   link.ShortCode,
			Visits:      link.Visits,
			CreatedAt:   link.CreatedAt,
		})
	}
	return c.JSON(response)
}

func DeleteLink(c *fiber.Ctx) error {
	id := cast.ToUint(c.Params("id"))
	userId := cast.ToUint(c.Locals(constants.USER_ID_KEY))

	var link models.Link
	database.DB.Where("id = ? AND user_id = ?", id, userId).First(&link)
	if link.ID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"error": "Link not found.",
			},
		)
	}
	result := database.DB.Delete(&link) // soft delete (data remains in db but is considered deleted)
	// result := database.DB.Unscoped().Delete(&link) // hard delete (data is removed from db)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"error": "Faild to delete link.",
			},
		)
	}

	return c.JSON(fiber.Map{
		"message": "link deleted successfuly",
	})
}
