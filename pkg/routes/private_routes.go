package routes

import (
	"admin-server/admin/controllers"

	"github.com/gofiber/fiber/v2"
)

func PrivateRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/admin")

	// Routes for POST method:
	route.Post("/notice", controllers.CreateNotice)       // post one notice
	route.Put("/notice", controllers.UpdateNotice)        // update one notice
	route.Delete("/notice/:id", controllers.DeleteNotice) // delete one notice by ID

	route.Post("/email", controllers.CreateEmail)
	route.Put("/email", controllers.UpdateEmail)
	route.Delete("/email/:id", controllers.DeleteEmail)
}
