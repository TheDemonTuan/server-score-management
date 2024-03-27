package router

import (
	"github.com/gofiber/fiber/v2"
	"qldiemsv/controllers"
)

func assignmentsRouter(r fiber.Router) {
	assignmentsRoute := r.Group("assignments")

	assignmentsRoute.Add("GET", "", controllers.AssignmentGetAll)
	//assignmentsRoute.Add("GET", ":id", controllers.AssignmentGetById)
	assignmentsRoute.Add("POST", "", controllers.AssignmentCreate)
	assignmentsRoute.Add("PUT", ":id", controllers.AssignmentUpdateById)
	//assignmentsRoute.Add("DELETE", "", controllers.AssignmentDeleteAll)
	//assignmentsRoute.Add("DELETE", "list", controllers.AssignmentDeleteByListId)
	assignmentsRoute.Add("DELETE", ":id", controllers.AssignmentDeleteById)
}
