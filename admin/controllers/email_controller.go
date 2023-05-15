package controllers

import (
	"admin-server/admin/models"
	"admin-server/platform/database"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetEmails godoc
// @Tags Email
// @Summary Get emails
// @Description Get emails
// @Accept json
// @Produce json
// @Success 200 {object} models.Email
// @Router /admin/emails [get]
func GetEmails(c *fiber.Ctx) error {
	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	emails, err := db.GetEmails()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":  true,
			"msg":    "Emails were not found",
			"count":  0,
			"emails": nil,
		})
	}

	return c.JSON(fiber.Map{
		"error":  false,
		"msg":    nil,
		"count":  len(emails),
		"emails": emails,
	})

}

// GetEmail godoc
// @Tags Email
// @Summary Get email
// @Description Get email
// @Accept json
// @Produce json
// @Param company_id path string true "Company ID"
// @Success 200 {object} models.Email
// @Router /admin/email/{company_id} [get]
func GetEmailsByCompanyName(c *fiber.Ctx) error {
	eid := models.CompanyIdDto{}
	id, err := uuid.Parse(c.Params("company_id"))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	eid.SetCompanyIdDto(id)

	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	emails, err := db.GetEmailsByCompnayId(&eid)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "email with the given ID is not found",
			"email": nil,
		})
	}

	return c.JSON(fiber.Map{
		"error":  false,
		"msg":    nil,
		"emails": emails,
	})
}

// @Description Create a new email.
// @Summary create a new email
// @Tags Email
// @Accept json
// @Produce json
// @Param email body string true "Email"
// @Success 200 {object} models.Email
// @Router /admin/email [post]
func CreateEmail(c *fiber.Ctx) error {
	email := &models.EmailUploadDto{}

	if err := c.BodyParser(email); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if err := db.CreateEmail(email.SetEmailUploadDtoToEntity()); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error":  false,
		"msg":    nil,
		"notice": email,
	})
}

// @Description Update email.
// @Summary update email
// @Tags Email
// @Accept json
// @Produce json
// @Param email body string true "Email"
// @Success 201 {string} status "ok"
// @Router /admin/email [put]
func UpdateEmail(c *fiber.Ctx) error {
	email := &models.EmailUpdateDto{}

	if err := c.BodyParser(email); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if err := db.UpdateEmail(email.SetEmailUpdateDtoToEntity()); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusCreated)
}

// DeleteEmail func for deletes email by given ID.
// @Description Delete email by given ID.
// @Summary delete email by given ID
// @Tags Email
// @Accept json
// @Produce json
// @Param id path int true "Email ID"
// @Success 204 {string} status "ok"
// @Router /admin/email/{id} [delete]
func DeleteEmail(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	validId := models.EmailValidIdDto{}
	validId.SetEmailValidDto(id)

	if err := db.DeleteEmail(&validId); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)

}
