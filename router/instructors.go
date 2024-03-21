package router

import (
	"github.com/gofiber/fiber/v2"
	"qldiemsv/controllers"
)

func instructorsRouter(r fiber.Router) {
	instructorsRoute := r.Group("instructors")

	instructorsRoute.Add("GET", "", controllers.InstructorGetAll)
	instructorsRoute.Add("GET", ":id", controllers.InstructorGetById)
	instructorsRoute.Add("POST", "", controllers.InstructorCreate)
	instructorsRoute.Add("PUT", ":id", controllers.InstructorUpdateById)
	instructorsRoute.Add("DELETE", "", controllers.InstructorDeleteAll)
	instructorsRoute.Add("DELETE", "list", controllers.InstructorDeleteByListId)
	instructorsRoute.Add("DELETE", ":id", controllers.InstructorDeleteById)
}
