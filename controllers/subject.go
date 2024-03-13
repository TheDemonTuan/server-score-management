package controllers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"math/rand"
	"qldiemsv/common"
	"qldiemsv/models/entity"
	"qldiemsv/models/req"
	"strconv"
)

func GenerateSubjectID() string {
	id := ""
	for i := 0; i < 6; i++ {
		id += strconv.Itoa(rand.Intn(10))
	}
	return "CS" + id
}

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

func SubjectCreate(c *fiber.Ctx) error {
	bodyData, err := common.Validator[req.SubjectCreate](c)

	if err != nil || bodyData == nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if totalPercentage := bodyData.ProcessPercentage + bodyData.MidtermPercentage + bodyData.FinalPercentage; totalPercentage != 100 {
		return fiber.NewError(fiber.StatusBadRequest, "Tổng phần trăm phải là 100")
	}

	newSubject := entity.Subject{
		ID:                GenerateSubjectID(),
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

	totalPercentage := bodyData.ProcessPercentage + bodyData.MidtermPercentage + bodyData.FinalPercentage
	if totalPercentage != 100 {
		return fiber.NewError(fiber.StatusBadRequest, "Tổng phần trăm phải là 100")
	}

	if bodyData.ProcessPercentage < 0 || bodyData.ProcessPercentage > 100 ||
		bodyData.MidtermPercentage < 0 || bodyData.MidtermPercentage > 100 ||
		bodyData.FinalPercentage < 0 || bodyData.FinalPercentage > 100 {
		return fiber.NewError(fiber.StatusBadRequest, "Phần trăm phải từ 0 đến 100")
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

// [DELETE] /api/subject/:id
func SubjectDelete(c *fiber.Ctx) error {
	id := c.Params("id")
	var subject entity.Subject

	if err := common.DBConn.First(&subject, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusBadRequest, "Không tìm thấy môn học")
		}
		return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi truy vấn cơ sở dữ liệu")
	}

	if err := common.DBConn.Delete(&subject).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi xóa môn học")
	}

	return c.JSON(common.NewResponse(fiber.StatusOK, "Success", nil))
}
