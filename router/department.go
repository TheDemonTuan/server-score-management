package router

import (
	"github.com/gofiber/fiber/v2"
	"qldiemsv/controllers"
)

func departmentRouter(r fiber.Router) {
	departmentRouter := r.Group("department")

	// Định nghĩa các route con của /api/department
	departmentRouter.Get("", controllers.DepartmentList)
	departmentRouter.Post("", controllers.DepartmentCreate)
}
