package handler

import (
	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"message": "ok",
	})
}

func GetAllUsers(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"message": "ok",
	})
}

func GetUser(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"message": "ok",
	})
}

func UpdateUser(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"message": "ok",
	})
}

func DeleteUser(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"message": "ok",
	})
}
