package controllers

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"math/rand"
	"qldiemsv/common"
	"qldiemsv/models/entity"
	"qldiemsv/models/req"
	"strconv"
	"time"
)

func generateTeacherID(departmentID int8) string {

	idPrefix := "GV"

	departmentCode := strconv.Itoa(int(departmentID))

	if len(departmentCode) == 1 {
		departmentCode = "0" + departmentCode
	}

	randomNumbers := fmt.Sprintf("%06d", rand.Intn(10000))

	teacherID := idPrefix + departmentCode + randomNumbers

	return teacherID
}

// [GET] /api/teacher
func TeacherList(c *fiber.Ctx) error {
	var teachers []entity.Teacher

	if err := common.DBConn.Preload("Subjects").Find(&teachers).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi truy vấn cơ sở dữ liệu")

	}
	return c.JSON(common.NewResponse(fiber.StatusOK, "Success", teachers))
}

// [GET] /api/teacher/:id
func TeacherGetById(c *fiber.Ctx) error {
	id := c.Params("id")
	var teacher entity.Teacher

	if err := common.DBConn.Preload("Subjects").First(&teacher, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusBadRequest, "Không tìm thấy giáo viên")
		} else {
			return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi truy vấn cơ sở dữ liệu")
		}
	}
	return c.JSON(common.NewResponse(fiber.StatusOK, "Success", teacher))
}

// [POST] /api/teacher
func TeacherCreate(c *fiber.Ctx) error {
	bodyData, err := common.Validator[req.TeacherCreate](c)

	if err != nil || bodyData == nil {

		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	ageLimit := 22
	today := time.Now()
	minBirthday := today.AddDate(-ageLimit, 0, 0)
	if bodyData.BirthDay.After(minBirthday) {
		return fiber.NewError(fiber.StatusBadRequest, "Tuổi giáo viên phải lớn hơn 22")

	}
	newTeacher := entity.Teacher{
		ID:           generateTeacherID(bodyData.DepartmentID),
		Name:         bodyData.Name,
		Email:        bodyData.Email,
		Phone:        bodyData.Phone,
		DepartmentID: bodyData.DepartmentID,
		BirthDay:     bodyData.BirthDay,
		Degree:       bodyData.Degree,
		Address:      bodyData.Address,
	}

	if err := common.DBConn.Create(&newTeacher).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi tạo giáo viên")
	}

	return c.JSON(common.NewResponse(fiber.StatusOK, "Success", newTeacher))

}

// [PUT] /api/teacher/:id
func TeacherUpdate(c *fiber.Ctx) error {
	bodyData, err := common.Validator[req.TeacherUpdate](c)

	if err != nil || bodyData == nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var teacher entity.Teacher

	if err := common.DBConn.First(&teacher, "id = ?", c.Params("id")).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusBadRequest, "Không tìm thấy giáo viên")
		} else {
			return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi truy vấn cơ sở dữ liệu")
		}
	}

	ageLimit := 22
	today := time.Now()
	minBirthday := today.AddDate(-ageLimit, 0, 0)
	if bodyData.BirthDay.After(minBirthday) {
		return fiber.NewError(fiber.StatusBadRequest, "Tuổi giáo viên phải lớn hơn 22")
	}

	teacher.Name = bodyData.Name
	teacher.Email = bodyData.Email
	teacher.Phone = bodyData.Phone
	teacher.DepartmentID = bodyData.DepartmentID
	teacher.BirthDay = bodyData.BirthDay
	teacher.Degree = bodyData.Degree
	teacher.Address = bodyData.Address

	if err := common.DBConn.Save(&teacher).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi cập nhật giáo viên")
	}

	return c.JSON(common.NewResponse(fiber.StatusOK, "Success", teacher))
}

// [DELETE] /api/teacher/:id
func TeacherDelete(c *fiber.Ctx) error {
	id := c.Params("id")
	var teacher entity.Teacher
	if err := common.DBConn.First(&teacher, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusBadRequest, "Không tìm thấy giáo viên")
		} else if errors.Is(err, gorm.ErrForeignKeyViolated) {
			return fiber.NewError(fiber.StatusBadRequest, "Không thể xóa giáo viên này vì có môn học thuộc giáo viên này")
		} else {
			return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi truy vấn cơ sở dữ liệu")
		}
	}

	if err := common.DBConn.Delete(&teacher).Error; err != nil {

		return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi xóa giáo viên")
	}

	return c.JSON(common.NewResponse(fiber.StatusOK, "Success", nil))

}
