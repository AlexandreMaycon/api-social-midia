package router

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func StartServer(port string) {
	app := fiber.New()

	UserRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
