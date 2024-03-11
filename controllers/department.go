package controllers

import (
	"github.com/gofiber/fiber/v2"
	"qldiemsv/common"
	"qldiemsv/models/entity"
	"qldiemsv/models/request"
)

// [GET] /api/department
func DepartmentList(c *fiber.Ctx) error {
	var departments []entity.Department

	result := common.DBConn.Find(&departments)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi truy vấn cơ sở dữ liệu")
	}

	return c.JSON(common.NewResponse(
		fiber.StatusOK,
		"Success",
		departments))
}

// [POST] /api/department
func DepartmentCreate(c *fiber.Ctx) error {

	department := new(request.DepartmentCreateRequest)

	if err := c.BodyParser(&department); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Lỗi khi parse request")
	}
	//
	newDepartment := entity.Department{
		ID:   department.ID,
		Name: department.Name,
	}

	errCreateDepartment := common.DBConn.Create(&newDepartment).Error
	if errCreateDepartment != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi tạo phòng ban")
	}

	return c.JSON(common.NewResponse(fiber.StatusOK, "Tạo phòng ban thành công", newDepartment))
}
