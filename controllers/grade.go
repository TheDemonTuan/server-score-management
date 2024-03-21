package controllers

import (
	"github.com/gofiber/fiber/v2"
	"qldiemsv/common"
	"qldiemsv/models/entity"
	"qldiemsv/models/req"
)

func validateGradeData(bodyData *req.GradeRequest) error {

	if err := common.DBConn.First(&entity.Subject{}, "id = ?", bodyData.SubjectID).Error; err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Không tìm thấy môn học")
	}

	if err := common.DBConn.First(&entity.Student{}, "id = ?", bodyData.StudentID).Error; err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Không tìm thấy sinh viên")
	}

	if err := common.DBConn.First(&entity.Instructor{}, "id = ?", bodyData.ByInstructorID).Error; err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Không tìm thấy giảng viên")
	}

	var instructor entity.Instructor
	if err := common.DBConn.First(&instructor, "id = ? AND id IN (SELECT instructor_id FROM assignments WHERE subject_id = ?)", bodyData.ByInstructorID, bodyData.SubjectID).Error; err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Giảng viên không dạy môn học này")
	}

	var student entity.Student
	if err := common.DBConn.
		Joins("JOIN subjects ON subjects.department_id = students.department_id").
		Where("students.id = ? AND subjects.id = ?", bodyData.StudentID, bodyData.SubjectID).
		First(&student).
		Error; err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Sinh viên không học môn học này")
	}
	return nil
}

// [GET] /api/grades
func GradeGetList(c *fiber.Ctx) error {
	var grades []entity.Grade

	if err := common.DBConn.Find(&grades).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi truy vấn cơ sở dữ liệu")
	}

	return c.JSON(common.NewResponse(
		fiber.StatusOK,
		"Success",
		grades))
}

// [POST] /api/grades
func GradeCreate(c *fiber.Ctx) error {
	bodyData, err := common.Validator[req.GradeRequest](c)

	if err != nil || bodyData == nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := validateGradeData(bodyData); err != nil {
		return err
	}

	newGrade := entity.Grade{
		ProcessScore:   bodyData.ProcessScore,
		MidtermScore:   bodyData.MidtermScore,
		FinalScore:     bodyData.FinalScore,
		SubjectID:      bodyData.SubjectID,
		StudentID:      bodyData.StudentID,
		ByInstructorID: bodyData.ByInstructorID,
	}

	if err := common.DBConn.Create(&newGrade).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi tạo điểm")
	}

	return c.JSON(common.NewResponse(fiber.StatusOK, "Success", newGrade))
}

// [GET] /api/grades/:id
func GradeGetById(c *fiber.Ctx) error {
	id := c.Params("id")
	var grade entity.Grade

	if err := common.DBConn.First(&grade, "id = ?", id).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi truy vấn cơ sở dữ liệu")
	}

	return c.JSON(common.NewResponse(fiber.StatusOK, "Success", grade))
}

// [PUT] /api/grades/:id
func GradeUpdateById(c *fiber.Ctx) error {
	id := c.Params("id")
	bodyData, err := common.Validator[req.GradeRequest](c)

	if err != nil || bodyData == nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var grade entity.Grade

	if err := common.DBConn.First(&grade, "id = ?", id).Error; err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Không tìm thấy điểm")
	}

	if err := validateGradeData(bodyData); err != nil {
		return err
	}

	grade.ProcessScore = bodyData.ProcessScore
	grade.MidtermScore = bodyData.MidtermScore
	grade.FinalScore = bodyData.FinalScore
	grade.SubjectID = bodyData.SubjectID
	grade.StudentID = bodyData.StudentID
	grade.ByInstructorID = bodyData.ByInstructorID

	if err := common.DBConn.Save(&grade).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi cập nhật điểm")
	}

	return c.JSON(common.NewResponse(fiber.StatusOK, "Success", grade))
}
