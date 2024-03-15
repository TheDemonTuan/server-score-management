package controllers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"qldiemsv/common"
	"qldiemsv/models/entity"
	"qldiemsv/models/req"
)

// [GET] /api/department
func DepartmentList(c *fiber.Ctx) error {
	// Khai báo một mảng chứa các phòng ban
	var departments []entity.Department

	// Lấy danh sách các phòng ban từ cơ sở dữ liệu
	if err := common.DBConn.Preload("Teachers").Preload("Subjects").Preload("Classes").Preload("Students").Find(&departments).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi truy vấn cơ sở dữ liệu")
	}
	// Trả về danh sách các phòng ban dưới dạng JSON
	return c.JSON(common.NewResponse(
		fiber.StatusOK,
		"Success",
		departments))
}

// [POST] /api/department
func DepartmentCreate(c *fiber.Ctx) error {
	bodyData, err := common.Validator[req.DepartmentCreate](c)

	if err != nil || bodyData == nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	newDepartment := entity.Department{
		Name: bodyData.Name,
	}

	if err := common.DBConn.Create(&newDepartment).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi tạo khoa")
	}

	return c.JSON(common.NewResponse(fiber.StatusOK, "Success", newDepartment))
}

// [GET] /api/department/:id
func DepartmentGetById(c *fiber.Ctx) error {
	id := c.Params("id")
	var department entity.Department

	if err := common.DBConn.Preload("Teachers").
		Preload("Subjects").Preload("Classes").Preload("Students").
		First(&department, "id = ?", id).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusBadRequest, "Không tìm thấy khoa")
		}

		return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi truy vấn cơ sở dữ liệu")
	}

	return c.JSON(common.NewResponse(fiber.StatusOK, "Success", department))
}

// [PUT] /api/department/:id
func DepartmentUpdate(c *fiber.Ctx) error {
	bodyData, err := common.Validator[req.DepartmentUpdate](c)

	if err != nil || bodyData == nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var department entity.Department

	id := c.Params("id")
	if err := common.DBConn.First(&department, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusBadRequest, "Không tìm thấy khoa")
		}
		return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi truy vấn cơ sở dữ liệu")
	}

	department.Name = bodyData.Name

	if err := common.DBConn.Save(&department).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi cập nhật khoa")
	}
	return c.JSON(common.NewResponse(fiber.StatusOK, "Success", department))
}

// [DELETE] /api/department/:id
func DepartmentDelete(c *fiber.Ctx) error {
	id := c.Params("id")

	var department entity.Department

	if err := common.DBConn.First(&department, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusBadRequest, "Không tìm thấy khoa")
		}

		return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi truy vấn cơ sở dữ liệu")
	}

	if err := common.DBConn.Delete(&department).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi xóa khoa")
	}

	return c.JSON(common.NewResponse(fiber.StatusOK, "Success", nil))
}
