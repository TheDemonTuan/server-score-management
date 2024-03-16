package router

import (
	"github.com/gofiber/fiber/v2"
	"qldiemsv/controllers"
)

func departmentsRouter(r fiber.Router) {
	departmentsRoute := r.Group("departments")

	departmentsRoute.Get("", controllers.DepartmentGetList)
	departmentsRoute.Get("/:id", controllers.DepartmentGetById)
	departmentsRoute.Post("", controllers.DepartmentCreate)
	departmentsRoute.Put("/:id", controllers.DepartmentUpdateById)
	departmentsRoute.Delete("/:id", controllers.DepartmentDeleteById)
}
