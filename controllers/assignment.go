package controllers

import (
	"github.com/gofiber/fiber/v2"
	"qldiemsv/common"
	"qldiemsv/models/entity"
	"qldiemsv/models/req"
)

// [GET] /api/assignments
func AssignmentGetList(c *fiber.Ctx) error {
	var assignments []entity.Assignment
	if err := common.DBConn.Find(&assignments).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi truy vấn cơ sở dữ liệu")
	}
	return c.JSON(common.NewResponse(
		fiber.StatusOK,
		"Success",
		assignments))
}

// [POST] /api/assignments
func AssignmentCreate(c *fiber.Ctx) error {
	bodyData, err := common.Validator[req.AssignmentCreate](c)
	if err != nil || bodyData == nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	var subject entity.Subject
	if err := common.DBConn.First(&subject, "id = ?", bodyData.SubjectID).Error; err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Không tìm thấy môn học")
	}

	var instructor entity.Instructor
	if err := common.DBConn.First(&instructor, "id = ?", bodyData.InstructorID).Error; err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Không tìm thấy giảng viên")
	}
	if instructor.DepartmentID != subject.DepartmentID {
		return fiber.NewError(fiber.StatusBadRequest, "Giảng viên không thuộc khoa của môn học")
	}

	newAssignment := entity.Assignment{
		SubjectID:    bodyData.SubjectID,
		InstructorID: bodyData.InstructorID,
	}

	if err := common.DBConn.First(&newAssignment, "subject_id = ? AND instructor_id = ?", bodyData.SubjectID, bodyData.InstructorID).Error; err == nil {
		return fiber.NewError(fiber.StatusBadRequest, "Phân công đã tồn tại")
	}

	if err := common.DBConn.Create(&newAssignment).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi tạo phân công")
	}
	return c.JSON(common.NewResponse(fiber.StatusOK, "Success", newAssignment))
}

// [GET] /api/assignments/:id
func AssignmentGetById(c *fiber.Ctx) error {
	id := c.Params("id")
	var assignment entity.Assignment
	if err := common.DBConn.First(&assignment, "id = ?", id).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi truy vấn cơ sở dữ liệu")
	}
	return c.JSON(common.NewResponse(fiber.StatusOK, "Success", assignment))
}

// [PUT] /api/assignments/:id
func AssignmentUpdateById(c *fiber.Ctx) error {
	id := c.Params("id")
	bodyData, err := common.Validator[req.AssignmentUpdate](c)
	if err != nil || bodyData == nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	var assignment entity.Assignment
	if err := common.DBConn.First(&assignment, "id = ?", id).Error; err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Không tìm thấy phân công")

	}
	var subject entity.Subject
	if err := common.DBConn.First(&subject, "id = ?", bodyData.SubjectID).Error; err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Không tìm thấy môn học")
	}

	var instructor entity.Instructor
	if err := common.DBConn.First(&instructor, "id = ?", bodyData.InstructorID).Error; err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Không tìm thấy giảng viên")
	}

	if instructor.DepartmentID != subject.DepartmentID {
		return fiber.NewError(fiber.StatusBadRequest, "Giảng viên không thuộc khoa của môn học")
	}

	assignment.SubjectID = bodyData.SubjectID
	assignment.InstructorID = bodyData.InstructorID

	if err := common.DBConn.First(&assignment, "subject_id = ? AND instructor_id = ?", bodyData.SubjectID, bodyData.InstructorID).Error; err == nil {
		return fiber.NewError(fiber.StatusBadRequest, "Phân công đã tồn tại")
	}
	if err := common.DBConn.Save(&assignment).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi cập nhật phân công")
	}
	return c.JSON(common.NewResponse(fiber.StatusOK, "Success", assignment))
}

// [DELETE] /api/assignments/:subjectID/:instructorID
func AssignmentDeleteById(c *fiber.Ctx) error {
	subjectID := c.Params("subjectID")
	instructorID := c.Params("instructorID")
	var assignment entity.Assignment
	if err := common.DBConn.First(&assignment, "subject_id = ? AND instructor_id = ?", subjectID, instructorID).Error; err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Không tìm thấy phân công")
	}
	if err := common.DBConn.Delete(&assignment).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi xóa phân công")
	}
	return c.JSON(common.NewResponse(fiber.StatusOK, "Success", nil))
}

// [DELETE] /api/assignments
func AssignmentDeleteAll(c *fiber.Ctx) error {
	if err := common.DBConn.Where("1 = 1").Delete(&entity.Assignment{}).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi xóa tất cả phân công")
	}
	return c.JSON(common.NewResponse(fiber.StatusOK, "Success", nil))
}

// [DELETE] /api/assignments/list
func AssignmentDeleteByListId(c *fiber.Ctx) error {
	var subjectIDs []string
	var instructorIDs []string
	if err := c.BodyParser(&subjectIDs); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Danh sách ID không hợp lệ")
	}
	if err := c.BodyParser(&instructorIDs); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Danh sách ID không hợp lệ")
	}
	if err := common.DBConn.Where("subject_id IN ? AND instructor_id IN ?", subjectIDs, instructorIDs).Delete(&entity.Assignment{}).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Lỗi khi xóa phân công")
	}
	return c.JSON(common.NewResponse(fiber.StatusOK, "Success", nil))
}
