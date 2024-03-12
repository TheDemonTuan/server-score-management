package controllers

import (
	"github.com/gofiber/fiber/v2"
	"qldiemsv/common"
	"qldiemsv/models/entity"
)

// [GET] /api/teacher
func TeacherList(c *fiber.Ctx) error {
	var teachers []entity.Teacher

	result := common.DBConn.Find(&teachers)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(common.NewResponse(
			fiber.StatusInternalServerError, "Lỗi khi truy vấn cơ sở dữ liệu", nil))
	}

	return c.JSON(common.NewResponse(
		fiber.StatusOK,
		"Success",
		teachers))
}

// [GET] /api/teacher/:id

// [PUT] /api/teacher/:id

// [DELETE] /api/teacher/:id
