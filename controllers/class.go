package controllers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"qldiemsv/common"
	"qldiemsv/models/entity"
	"qldiemsv/models/req"
)

// [GET] /api/classes
func ClassGetAll(c *fiber.Ctx) error {
	var classes []entity.Class

	if err := common.DBConn.Preload("Students").Find(&classes).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi truy vấn cơ sở dữ liệu")
	}

	return c.JSON(common.NewResponse(
		fiber.StatusOK,
		"Success",
		classes))
}

// [GET] /api/classes/:id
func ClassGetById(c *fiber.Ctx) error {
	classId := c.Params("id")
	var class entity.Class

	if err := common.DBConn.Preload("Students").First(&class, "id = ?", classId).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi truy vấn cơ sở dữ liệu")
	}

	return c.JSON(common.NewResponse(fiber.StatusOK, "Success", class))
}

// [POST] /api/classes
func ClassCreate(c *fiber.Ctx) error {
	bodyData, err := common.Validator[req.ClassCreate](c)

	if err != nil || bodyData == nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	newClass := entity.Class{
		Name:             bodyData.Name,
		MaxStudents:      bodyData.MaxStudents,
		DepartmentID:     bodyData.DepartmentID,
		HostInstructorID: bodyData.HostInstructorID,
	}

	if err := common.DBConn.Create(&newClass).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi tạo lớp")
	}

	return c.JSON(common.NewResponse(fiber.StatusOK, "Success", newClass))
}

// [PUT] /api/classes/:id
func ClassUpdateById(c *fiber.Ctx) error {
	classId := c.Params("id")
	bodyData, err := common.Validator[req.ClassUpdateById](c)

	if err != nil || bodyData == nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var class entity.Class

	if err := common.DBConn.First(&class, "id = ?", classId).Error; err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Không tìm thấy lớp")
	}

	class.Name = bodyData.Name
	class.MaxStudents = bodyData.MaxStudents
	class.DepartmentID = bodyData.DepartmentID
	class.HostInstructorID = bodyData.HostInstructorID

	if err := common.DBConn.Save(&class).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi cập nhật lớp")
	}

	return c.JSON(common.NewResponse(fiber.StatusOK, "Success", class))
}

// [DELETE] /api/classes/:id
func ClassDeleteById(c *fiber.Ctx) error {
	classId := c.Params("id")
	var class entity.Class

	if err := common.DBConn.First(&class, "id = ?", classId).Error; err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Không tìm thấy lớp")
	}

	if err := common.DBConn.Delete(&class).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi xóa lớp")
	}

	return c.JSON(common.NewResponse(fiber.StatusOK, "Success", nil))
}

// [DELETE] /api/classes
func ClassDeleteAll(c *fiber.Ctx) error {
	if err := common.DBConn.Where("1 = 1").Delete(&entity.Class{}).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi xóa tất cả lớp")
	}

	return c.JSON(common.NewResponse(fiber.StatusOK, "Success", nil))
}

// [DELETE] /api/classes/list
func ClassDeleteByListId(c *fiber.Ctx) error {
	bodyData, err := common.Validator[req.ClassDeleteByListId](c)

	if err != nil || bodyData == nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := common.DBConn.Where("id IN ?", bodyData.ListId).Delete(&entity.Class{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusBadRequest, "Không tìm thấy lớp")
		}
		return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi xóa nhiều lớp")
	}
	return c.JSON(common.NewResponse(fiber.StatusOK, "Success", nil))
}
