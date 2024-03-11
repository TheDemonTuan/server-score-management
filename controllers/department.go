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
	if err := common.DBConn.Find(&departments).Error; err != nil {
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
	//Hàm validate t viết sẵn chỉ cần định nghĩa validate ở bên type xong gọi như bên dưới thay cais req.???? tuỳ theo tên của req
	bodyData, err := common.Validator[req.DepartmentCreate](c)

	// Nếu dữ liệu không hợp lệ hoặc không có dữ liệu thì trả về lỗi 400
	if err != nil || bodyData == nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	newDepartment := entity.Department{
		ID:   bodyData.ID,
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

	if err := common.DBConn.First(&department, "id = ?", id).Error; err != nil {
		// đây là check coi lỗi trả ve có phải là not found hay không đọc document gorm chỗ error handling thì đây là do người dùng truyền sai id nên không có dữ liệu nên status bad req mới đúng
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusBadRequest, "Không tìm thấy khoa")
		} else {
			return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi truy vấn cơ sở dữ liệu")
		} // đây là các lỗi chưa biết thì trả về lỗi 500 do chưa biết lỗi từ đau mà ra
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
		} else {
			return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi truy vấn cơ sở dữ liệu")
		}
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
		} else {
			return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi truy vấn cơ sở dữ liệu")
		}
	}

	if err := common.DBConn.Delete(&department).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi xóa khoa")
	}

	return c.JSON(common.NewResponse(fiber.StatusOK, "Success", nil))
}
