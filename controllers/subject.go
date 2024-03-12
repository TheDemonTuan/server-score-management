package controllers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"qldiemsv/common"
	"qldiemsv/models/entity"
	"qldiemsv/models/req"
)

// [GET] /api/subject
func SubjectList(c *fiber.Ctx) error {

	var subjects []entity.Subject

	if err := common.DBConn.Find(&subjects).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi truy vấn cơ sở dữ liệu")
	}

	return c.JSON(common.NewResponse(
		fiber.StatusOK,
		"Success",
		subjects))
}

// [POST] /api/subject
func SubjectCreate(c *fiber.Ctx) error {
	bodyData, err := common.Validator[req.SubjectCreate](c)

	if err != nil || bodyData == nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	newSubject := entity.Subject{
		ID:                bodyData.ID,
		Name:              bodyData.Name,
		Credits:           bodyData.Credits,
		ProcessPercentage: bodyData.ProcessPercentage,
		MidtermPercentage: bodyData.MidtermPercentage,
		FinalPercentage:   bodyData.FinalPercentage,
		DepartmentID:      bodyData.DepartmentID,
	}

	if err := common.DBConn.Create(&newSubject).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi tạo môn học")
	}
	return c.JSON(common.NewResponse(fiber.StatusOK, "Success", newSubject))
}

// [GET] /api/subject/:id
func SubjectGetById(c *fiber.Ctx) error {
	id := c.Params("id")
	var subject entity.Subject

	if err := common.DBConn.First(&subject, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusBadRequest, "Không tìm thấy khoa")
		} else {
			return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi truy vấn cơ sở dữ liệu")
		}
	}
	return c.JSON(common.NewResponse(fiber.StatusOK, "Success", subject))
}

// [PUT] /api/subject/:id
func SubjectUpdate(c *fiber.Ctx) error {
	bodyData, err := common.Validator[req.SubjectUpdate](c)

	if err != nil || bodyData == nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var subject entity.Subject

	id := c.Params("id")
	if err := common.DBConn.First(&subject, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusBadRequest, "Không tìm thấy môn học")
		}
		return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi truy vấn cơ sở dữ liệu")

	}

	subject.Name = bodyData.Name
	subject.Credits = bodyData.Credits
	subject.ProcessPercentage = bodyData.ProcessPercentage
	subject.MidtermPercentage = bodyData.MidtermPercentage
	subject.FinalPercentage = bodyData.FinalPercentage
	subject.DepartmentID = bodyData.DepartmentID

	if err := common.DBConn.Save(&subject).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi cập nhật môn học")
	}

	return c.JSON(common.NewResponse(fiber.StatusOK, "Success", subject))
}
