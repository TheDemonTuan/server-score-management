package router

import (
	"github.com/gofiber/fiber/v2"
	"qldiemsv/controllers"
)

func departmentsRouter(r fiber.Router) {
	departmentsRoute := r.Group("departments")

	departmentsRoute.Add("GET", "", controllers.DepartmentGetAll)
	departmentsRoute.Add("GET", ":id", controllers.DepartmentGetById)
	departmentsRoute.Add("POST", "", controllers.DepartmentCreate)
	departmentsRoute.Add("PUT", ":id", controllers.DepartmentUpdateById)
	departmentsRoute.Add("DELETE", "", controllers.DepartmentDeleteAll)
	departmentsRoute.Add("DELETE", "list", controllers.DepartmentDeleteByListId)
	departmentsRoute.Add("DELETE", ":id", controllers.DepartmentDeleteById)
}
