package controllers

import (
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
	//
	newDepartment := entity.Department{
		ID:   department.ID,
		Name: department.Name,
	}

	errCreateDepartment := common.DBConn.Create(&newDepartment).Error
	if errCreateDepartment != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(common.NewResponse(
			fiber.StatusInternalServerError, "Lỗi khi tạo phòng ban", nil))
	}

	return c.JSON(common.NewResponse(fiber.StatusOK, "Success", newDepartment))
}
