package controllers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"qldiemsv/common"
	"qldiemsv/models/entity"
	"qldiemsv/models/req"
)

// [GET] /api/registrations
func RegistrationGetAll(c *fiber.Ctx) error {
	var registrations []entity.StudentRegistration

	if err := common.DBConn.Find(&registrations).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi truy vấn cơ sở dữ liệu")
	}

	return c.JSON(common.NewResponse(
		fiber.StatusOK,
		"Success",
		registrations))
}

// [POST] /api/registrations
func RegistrationCreate(c *fiber.Ctx) error {
	bodyData, err := common.Validator[req.RegistrationCreate](c)
	if err != nil || bodyData == nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var subject entity.Subject
	if err := common.DBConn.First(&subject, "id = ?", bodyData.SubjectID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusBadRequest, "Không tìm thấy môn học")
		}
		return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi truy vấn cơ sở dữ liệu")
	}

	var student entity.Student
	if err := common.DBConn.First(&student, "id = ?", bodyData.StudentID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusBadRequest, "Không tìm thấy sinh viên")
		}
		return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi truy vấn cơ sở dữ liệu")
	}

	if student.DepartmentID != subject.DepartmentID {
		return fiber.NewError(fiber.StatusBadRequest, "Sinh viên không thuộc khoa của môn học")
	}

	var registration entity.StudentRegistration
	if err := common.DBConn.First(&registration, "subject_id = ? AND student_id = ?", subject.ID, student.ID).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusBadRequest, "Lỗi khi truy vấn cơ sở dữ liệu")
		}
	}

	if registration.ID != 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Sinh viên đã đăng ký môn học này")
	}

	newRegistration := entity.StudentRegistration{
		SubjectID: bodyData.SubjectID,
		StudentID: bodyData.StudentID,
	}

	if err := common.DBConn.Create(&newRegistration).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi đăng ký môn học")
	}
	return c.JSON(common.NewResponse(fiber.StatusOK, "Success", newRegistration))
}
