package controllers

import (
	"github.com/gofiber/fiber/v2"
	"qldiemsv/common"
	"qldiemsv/models/entity"
	"qldiemsv/models/req"
)

// [GET] /api/classes
func ClassGetList(c *fiber.Ctx) error {
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
	id := c.Params("id")
	var class entity.Class

	if err := common.DBConn.Preload("Students").First(&class, "id = ?", id).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi truy vấn cơ sở dữ liệu")
	}

	return c.JSON(common.NewResponse(fiber.StatusOK, "Success", class))
}

// [POST] /api/classes
func ClassCreate(c *fiber.Ctx) error {
	bodyData, err := common.Validator[req.CreateClass](c)

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
func ClassUpdate(c *fiber.Ctx) error {
	id := c.Params("id")
	bodyData, err := common.Validator[req.UpdateClass](c)

	if err != nil || bodyData == nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var class entity.Class

	if err := common.DBConn.First(&class, "id = ?", id).Error; err != nil {
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

//// [DELETE] /api/classes/:id
//func ClassDelete(c *fiber.Ctx) error {
//	id := c.Params("id")
//	var class entity.Class
//
//	if err := common.DBConn.First(&class, "id = ?", id).Error; err != nil {
//		return fiber.NewError(fiber.StatusBadRequest, "Không tìm thấy lớp")
//	}
//
//	if err := common.DBConn.Delete(&class).Error; err != nil {
//		return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi xóa lớp")
//	}
//
//	return c.JSON(common.NewResponse(fiber.StatusOK, "Success", nil))
//}

// [DELETE] /api/classess/all
func ClassDeleteAll(c *fiber.Ctx) error {
	if err := common.DBConn.Where("1 = 1").Delete(&entity.Class{}).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi xóa tất cả lớp")
	}

	return c.JSON(common.NewResponse(fiber.StatusOK, "Success", nil))
}

// [DELETE] /api/classes
func ClassDeleteList(c *fiber.Ctx) error {
	var ids []string

	if err := c.BodyParser(&ids); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Danh sách ID không hợp lệ")
	}

	if err := common.DBConn.Where("id IN ?", ids).Delete(&entity.Class{}).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi xóa lớp")
	}
	return c.JSON(common.NewResponse(fiber.StatusOK, "Success", nil))
}
