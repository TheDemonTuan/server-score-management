package controllers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"qldiemsv/common"
	"qldiemsv/models/entity"
	"qldiemsv/models/req"
)

// [GET] /api/departments
func DepartmentGetAll(c *fiber.Ctx) error {
	isPreload := c.QueryBool("preload", true)
	selectFields := c.Query("select", "*")

	var departments []entity.Department

	if isPreload {
		if err := common.DBConn.Select(selectFields).Preload("Instructors").Preload("Subjects").Preload("Classes").Preload("Students").Find(&departments).Error; err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi truy vấn cơ sở dữ liệu")
		}
	} else {
		if err := common.DBConn.Select(selectFields).Find(&departments).Error; err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi truy vấn cơ sở dữ liệu")
		}
	}

	return c.JSON(common.NewResponse(
		fiber.StatusOK,
		"Success",
		departments))
}

// [POST] /api/departments
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

// [GET] /api/departments/:id
func DepartmentGetById(c *fiber.Ctx) error {
	departmentId := c.Params("id")
	isPreload := c.QueryBool("preload", true)
	selectFields := c.Query("select", "*")

	var department entity.Department

	if isPreload {
		if err := common.DBConn.Select(selectFields).Preload("Instructors").
			Preload("Subjects").Preload("Classes").Preload("Students").
			First(&department, "id = ?", departmentId).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return fiber.NewError(fiber.StatusBadRequest, "Không tìm thấy khoa")
			}
			return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi truy vấn cơ sở dữ liệu")
		}
	} else {
		if err := common.DBConn.Select(selectFields).First(&department, "id = ?", departmentId).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return fiber.NewError(fiber.StatusBadRequest, "Không tìm thấy khoa")
			}
			return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi truy vấn cơ sở dữ liệu")
		}

	}

	return c.JSON(common.NewResponse(fiber.StatusOK, "Success", department))
}

// [PUT] /api/departments/:id
func DepartmentUpdateById(c *fiber.Ctx) error {
	id, idErr := c.ParamsInt("id")

	if idErr != nil {
		return fiber.NewError(fiber.StatusBadRequest, "ID không hợp lệ")
	}

	bodyData, err := common.Validator[req.DepartmentUpdateById](c)

	if err != nil || bodyData == nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var department entity.Department

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

// [DELETE] /api/departments/:id
func DepartmentDeleteById(c *fiber.Ctx) error {
	id, idErr := c.ParamsInt("id")

	if idErr != nil {
		return fiber.NewError(fiber.StatusBadRequest, "ID không hợp lệ")
	}

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

// [DELETE] /api/departments/list
func DepartmentDeleteByListId(c *fiber.Ctx) error {
	bodyData, err := common.Validator[req.DepartmentDeleteByListId](c)

	if err != nil || bodyData == nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := common.DBConn.Where("id IN ?", bodyData.ListId).Delete(&entity.Department{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusBadRequest, "Không tìm thấy khoa")
		}
		return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi xóa nhiều khoa")
	}

	return c.JSON(common.NewResponse(fiber.StatusOK, "Success", nil))
}

// [DELETE] /api/departments
func DepartmentDeleteAll(c *fiber.Ctx) error {
	if err := common.DBConn.Where("1 = 1").Delete(&entity.Department{}).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi xóa tất cả khoa")
	}

	return c.JSON(common.NewResponse(fiber.StatusOK, "Success", nil))
}
