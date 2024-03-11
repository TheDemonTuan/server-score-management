package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"qldiemsv/common"
	"qldiemsv/models/entity"
	"qldiemsv/models/request"
)

// [GET] /api/department
func DepartmentList(c *fiber.Ctx) error {
	// Khai báo một mảng chứa các phòng ban
	var departments []entity.Department

	// Lấy danh sách các phòng ban từ cơ sở dữ liệu
	result := common.DBConn.Find(&departments)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(common.NewResponse(
			fiber.StatusInternalServerError, "Lỗi khi truy vấn cơ sở dữ liệu", nil))
	}
	// Trả về danh sách các phòng ban dưới dạng JSON
	return c.JSON(common.NewResponse(
		fiber.StatusOK,
		"Success",
		departments))
}

// [POST] /api/department
func DepartmentCreate(c *fiber.Ctx) error {

	department := new(request.DepartmentCreateRequest)

	if err := c.BodyParser(department); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewResponse(
			fiber.StatusBadRequest, "Lỗi khi parse request", nil))
	}

	validate := validator.New()
	errValidate := validate.Struct(department)
	if errValidate != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewResponse(
			fiber.StatusBadRequest, "Dữ liệu không hợp lệ", nil))
	}

	newDepartment := entity.Department{
		ID:   department.ID,
		Name: department.Name,
	}

	errCreateDepartment := common.DBConn.Create(&newDepartment).Error
	if errCreateDepartment != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(common.NewResponse(
			fiber.StatusInternalServerError, "Lỗi khi tạo khoa", nil))
	}

	return c.JSON(common.NewResponse(fiber.StatusOK, "Success", newDepartment))
}

// [GET] /api/department/:id
func DepartmentGetById(c *fiber.Ctx) error {
	id := c.Params("id")
	var department entity.Department
	err := common.DBConn.First(&department, "id = ?", id).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(common.NewResponse(
			fiber.StatusNotFound, "Không tìm thấy khoa", nil))
	}
	return c.JSON(common.NewResponse(fiber.StatusOK, "Success", department))
}

// [PUT] /api/department/:id
func DepartmentUpdate(c *fiber.Ctx) error {
	departmentRequest := new(request.DepartmentUpdateRequest)

	if err := c.BodyParser(departmentRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewResponse(
			fiber.StatusBadRequest, "Lỗi khi parse request", nil))
	}

	var department entity.Department
	id := c.Params("id")
	err := common.DBConn.First(&department, "id = ?", id).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(common.NewResponse(
			fiber.StatusNotFound, "Không tìm thấy khoa", nil))
	}

	if departmentRequest.Name != "" {
		department.Name = departmentRequest.Name
	}

	errUpdate := common.DBConn.Save(&department).Error
	if errUpdate != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(common.NewResponse(
			fiber.StatusInternalServerError, "Lỗi khi cập nhật khoa", nil))
	}

	return c.JSON(common.NewResponse(fiber.StatusOK, "Success", department))
}

// [DELETE] /api/department/:id
func DepartmentDelete(c *fiber.Ctx) error {
	id := c.Params("id")
	var department entity.Department
	err := common.DBConn.Debug().First(&department, "id = ?", id).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(common.NewResponse(
			fiber.StatusNotFound, "Không tìm thấy khoa", nil))
	}

	errDelete := common.DBConn.Debug().Delete(&department).Error
	if errDelete != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(common.NewResponse(
			fiber.StatusInternalServerError, "Lỗi khi xóa khoa", nil))
	}

	return c.JSON(common.NewResponse(fiber.StatusOK, "Success", nil))
}
