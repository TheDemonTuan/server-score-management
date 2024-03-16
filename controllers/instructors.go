package controllers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"qldiemsv/common"
	"qldiemsv/models/entity"
	"qldiemsv/models/req"
	"strconv"
	"time"
)

func generateInstructorID(departmentID uint) string {
	const maxLength = 10
	const idPrefix = "GV"
	departmentCode := strconv.Itoa(int(departmentID))

	return idPrefix + departmentCode + common.GenerateRandNum(maxLength-len(idPrefix)-len(departmentCode))
}

// [GET] /api/instructors
func InstructorGetList(c *fiber.Ctx) error {
	var instructors []entity.Instructor

	if err := common.DBConn.Preload("Classes").Preload("Grades").Preload("Assignments").Find(&instructors).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi truy vấn cơ sở dữ liệu")

	}

	return c.JSON(common.NewResponse(fiber.StatusOK, "Success", instructors))
}

// [GET] /api/instructors/:id
func InstructorGetById(c *fiber.Ctx) error {
	id := c.Params("id")
	var instructor entity.Instructor

	if err := common.DBConn.Preload("Subjects").First(&instructor, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusBadRequest, "Không tìm thấy giảng viên")
		}
		return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi truy vấn cơ sở dữ liệu")
	}

	return c.JSON(common.NewResponse(fiber.StatusOK, "Success", instructor))
}

// [POST] /api/instructors
func InstructorCreate(c *fiber.Ctx) error {
	bodyData, err := common.Validator[req.InstructorCreate](c)

	if err != nil || bodyData == nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	today := time.Now()
	if bodyData.BirthDay.After(today) {
		return fiber.NewError(fiber.StatusBadRequest, "Ngày sinh không hợp lệ")
	}

	var instructor entity.Instructor

	if err := common.DBConn.First(&instructor, "email = ? or phone = ?", bodyData.Email).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi truy vấn cơ sở dữ liệu")
		}
	}

	if instructor.ID != "" {
		return fiber.NewError(fiber.StatusBadRequest, "Email hoặc số điện thoại đã tồn tại")
	}

	newInstructor := entity.Instructor{
		ID:           generateInstructorID(bodyData.DepartmentID),
		FirstName:    bodyData.FirstName,
		LastName:     bodyData.LastName,
		Email:        bodyData.Email,
		Address:      bodyData.Address,
		Degree:       bodyData.Degree,
		BirthDay:     bodyData.BirthDay,
		Phone:        bodyData.Phone,
		Gender:       bodyData.Gender,
		DepartmentID: bodyData.DepartmentID,
	}

	if err := common.DBConn.Create(&newInstructor).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi tạo giảng viên")
	}

	return c.JSON(common.NewResponse(fiber.StatusOK, "Success", newInstructor))

}

// [PUT] /api/instructors/:id
func InstructorUpdateById(c *fiber.Ctx) error {
	bodyData, err := common.Validator[req.InstructorUpdateById](c)

	if err != nil || bodyData == nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	//Logic check
	today := time.Now()
	if bodyData.BirthDay.After(today) {
		return fiber.NewError(fiber.StatusBadRequest, "Ngày sinh không hợp lệ")
	}

	var instructor entity.Instructor

	if err := common.DBConn.First(&instructor, "email = ? or phone = ?", bodyData.Email).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi truy vấn cơ sở dữ liệu")
		}
	}

	if instructor.ID != "" {
		return fiber.NewError(fiber.StatusBadRequest, "Email hoặc số điện thoại đã tồn tại")
	}

	id := c.Params("id")
	if err := common.DBConn.First(&instructor, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusBadRequest, "Không tìm thấy giảng viên")
		}
		return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi truy vấn cơ sở dữ liệu")
	}
	//End logic check

	instructor.FirstName = bodyData.FirstName
	instructor.LastName = bodyData.LastName
	instructor.Email = bodyData.Email
	instructor.Address = bodyData.Address
	instructor.Degree = bodyData.Degree
	instructor.BirthDay = bodyData.BirthDay
	instructor.Phone = bodyData.Phone
	instructor.Gender = bodyData.Gender
	instructor.DepartmentID = bodyData.DepartmentID

	if err := common.DBConn.Save(&instructor).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi cập nhật giảng viên")
	}

	return c.JSON(common.NewResponse(fiber.StatusOK, "Success", instructor))
}

// [DELETE] /api/instructors/:id
func InstructorDeleteById(c *fiber.Ctx) error {
	id := c.Params("id")
	var instructor entity.Instructor

	if err := common.DBConn.First(&instructor, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusBadRequest, "Không tìm thấy giảng viên")
		}
		return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi truy vấn cơ sở dữ liệu")

	}

	if err := common.DBConn.Delete(&instructor).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi xóa giáo viên")
	}

	return c.JSON(common.NewResponse(fiber.StatusOK, "Success", nil))
}