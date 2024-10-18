package router

import (
	"github.com/AlexandreMaycon/api-social-midia.git/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	r := app.Group("/user")

	r.Post("/", handler.CreateUser)
	r.Get("/", handler.GetAllUsers)
	r.Put("/", handler.UpdateUser)
	r.Delete("/", handler.DeleteUser)
}
