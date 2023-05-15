package controllers

import (
	"admin-server/admin/models"
	"admin-server/platform/database"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetNotices godoc
// @Summary Get notices
// @Description Get notices
// @Accept json
// @Produce json
// @Success 200 {object} models.Notice
// @Router /admin/notices [get]
func GetNotices(c *fiber.Ctx) error {
	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	notices, err := db.GetNotices()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"msg":     "Notices were not found",
			"count":   0,
			"notices": nil,
		})
	}

	return c.JSON(fiber.Map{
		"error":   false,
		"msg":     nil,
		"count":   len(notices),
		"notices": notices,
	})
}

// GetNotice godoc
// @Summary Get a notice by ID
// @Description Get notice by ID
// @Accept json
// @Produce json
// @Param id path int true "Notice ID"
// @Success 200 {object} models.Notice
// @Router /admin/notice/{id} [get]
func GetNotice(c *fiber.Ctx) error {
	nid := models.NoticeIdDto{}
	id, err := strconv.Atoi(c.Params("id"))
	nid.SetNoticeId(id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
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

	// dto의 id값으로 조회
	notice, err := db.GetNotice(&nid)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":  true,
			"msg":    "notice with the given ID is not found",
			"notice": nil,
		})
	}

	return c.JSON(fiber.Map{
		"error":  false,
		"msg":    nil,
		"notice": notice,
	})
}

// @Description Create a new notice.
// @Summary create a new notice
// @Tags Notice
// @Accept json
// @Produce json
// @Param notice body string true "Notice"
// @Success 200 {object} models.Notice
// @Router /admin/notice [post]
func CreateNotice(c *fiber.Ctx) error {
	// jwt
	//now := time.Now().Unix()

	//claims, err := utils.ExtractTokenMetadata(c)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"error": true,
	// 		"msg":   err.Error(),
	// 	})
	// }

	// //expires := claims.Expires

	// if now > expires {
	// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
	// 		"error": true,
	// 		"msg":   "unauthorized, check expiration time of your token",
	// 	})
	// }

	notice := &models.NoticeUploadDto{}

	if err := c.BodyParser(notice); err != nil {
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

	if err := db.CreateNotice(notice.SetNoticeDtoToEntity()); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error":  false,
		"msg":    nil,
		"notice": notice,
	})

}

// UpdateNotice func for updates notice by given ID.
// @Description Update notice.
// @Summary update notice
// @Tags Notice
// @Accept json
// @Produce json
// @Param notice body string true "Notice"
// @Success 201 {string} status "ok"
// @Router /admin/notice [put]
func UpdateNotice(c *fiber.Ctx) error {
	notice := &models.Notice{}

	if err := c.BodyParser(notice); err != nil {
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

	foundedNotice, err := db.GetNotice(notice.SetNoticeIdDto())
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "Notice with this ID not found",
		})
	}

	// dto 생성
	nid := &models.NoticeIdDto{}
	nid.SetNoticeId(foundedNotice.NoticeId)

	notice.UpdatedAt = time.Now()

	// entity로 넣어주기
	if err := db.UpdatedNotice(nid.SetNoticeIdToEntity(), notice.SetUploadNoticeToEntity()); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusCreated)
}

// DeleteNotice func for deletes notice by given ID.
// @Description Delete notice by given ID.
// @Summary delete notice by given ID
// @Tags Notice
// @Accept json
// @Produce json
// @Param id path int true "Notice ID"
// @Success 204 {string} status "ok"
// @Router /admin/notice/{id} [delete]
func DeleteNotice(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	noticeDto := &models.NoticeDeleteDto{}
	noticeDto.NoticeId = id
	// 추후 jwt에서 꺼내는 로직으로 수정해야함
	uid, err := uuid.Parse("c266d4ca-4ea4-4d96-8fc1-4a494da16c6b")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	noticeDto.ProfileId = uid

	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	foundedNotice, err := db.GetNotice(noticeDto.SetNoticeIdDto())
	if err != nil {
		// Return status 404 and book not found error.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "book with this ID not found",
		})
	}

	if err := db.DeleteNotice(foundedNotice.SetDeleteNoticeToEntity()); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status 204 no content.
	return c.SendStatus(fiber.StatusNoContent)
}
