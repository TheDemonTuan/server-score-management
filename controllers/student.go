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
)

func generateStudentID(departmentID int8) string {

	idPrefix := "DH"

	departmentCode := strconv.Itoa(int(departmentID))

	if len(departmentCode) == 1 {
		departmentCode = "0" + departmentCode
	}

	randomNumbers := fmt.Sprintf("%06d", rand.Intn(10000))

	studentID := idPrefix + departmentCode + randomNumbers

	return studentID
}

// [GET] /api/student
func StudentList(c *fiber.Ctx) error {
	var students []entity.Student
	if err := common.DBConn.Preload("Transcripts").Find(&students).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi truy vấn cơ sở dữ liệu")
	}
	return c.JSON(common.NewResponse(
		fiber.StatusOK,
		"Success",
		students))
}

// [GET] /api/student/:id
// Hiển thị chi tiết sinh viên - Student - Subject - Transcript
func StudentGetById(c *fiber.Ctx) error {
	id := c.Params("id")
	var student entity.Student
	if err := common.DBConn.Preload("Transcripts").First(&student, "id = ?", id).Error; err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Không tìm thấy sinh viên")
	}

	var subjectIDs []string
	for _, transcript := range student.Transcripts {
		subjectIDs = append(subjectIDs, transcript.SubjectID)
	}

	var subjects []entity.Subject
	if err := common.DBConn.Preload("Transcripts").Find(&subjects, "id IN (?)", subjectIDs).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi truy vấn cơ sở dữ liệu")
	}

	var studentDetails struct {
		Student  entity.Student   `json:"Student"`
		Subjects []entity.Subject `json:"Subjects"`
	}

	studentDetails.Student = student
	studentDetails.Subjects = subjects

	return c.JSON(common.NewResponse(fiber.StatusOK, "Success", studentDetails))

}

// [POST] /api/student
func StudentCreate(c *fiber.Ctx) error {
	bodyData, err := common.Validator[req.StudentCreate](c)
	if err != nil || bodyData == nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	newStudent := entity.Student{
		ID:           generateStudentID(bodyData.DepartmentID),
		FirstName:    bodyData.FirstName,
		LastName:     bodyData.LastName,
		BirthDay:     bodyData.BirthDay,
		Gender:       bodyData.Gender,
		Email:        bodyData.Email,
		Phone:        bodyData.Phone,
		Address:      bodyData.Address,
		ClassID:      bodyData.ClassID,
		DepartmentID: bodyData.DepartmentID,
	}

	if err := common.DBConn.Create(&newStudent).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi tạo sinh viên")
	}
	return c.JSON(common.NewResponse(fiber.StatusOK, "Success", newStudent))
}

// [PUT] /api/student/:id
func StudentUpdate(c *fiber.Ctx) error {
	bodyData, err := common.Validator[req.StudentUpdate](c)
	if err != nil || bodyData == nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var student entity.Student
	id := c.Params("id")
	if err := common.DBConn.First(&student, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusBadRequest, "Không tìm thấy sinh viên")
		}
		return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi truy vấn cơ sở dữ liệu")
	}

	student.FirstName = bodyData.FirstName
	student.LastName = bodyData.LastName
	student.BirthDay = bodyData.BirthDay
	student.Gender = bodyData.Gender
	student.Email = bodyData.Email
	student.Phone = bodyData.Phone
	student.Address = bodyData.Address
	student.ClassID = bodyData.ClassID
	student.DepartmentID = bodyData.DepartmentID

	if err := common.DBConn.Save(&student).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi cập nhật sinh viên")
	}

	return c.JSON(common.NewResponse(fiber.StatusOK, "Success", student))

}

// [DELETE] /api/student/:id
func StudentDelete(c *fiber.Ctx) error {
	id := c.Params("id")
	var student entity.Student

	if err := common.DBConn.First(&student, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusBadRequest, "Không tìm thấy sinh viên")
		}
		return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi truy vấn cơ sở dữ liệu")
	}

	if err := common.DBConn.Delete(&student).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi xóa sinh viên")
	}

	return c.JSON(common.NewResponse(fiber.StatusOK, "Success", nil))
}
