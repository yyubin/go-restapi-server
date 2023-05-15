package routes

import (
	"admin-server/admin/controllers"

	"github.com/gofiber/fiber/v2"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/admin")

	// Routes for GET method:
	route.Get("/notices", controllers.GetNotices)   // get list of all notices
	route.Get("/notice/:id", controllers.GetNotice) // get one notice by ID

	route.Get("/emails", controllers.GetEmails)
	route.Get("/email/:company_id", controllers.GetEmailsByCompanyName)
}
