package handlers

import "github.com/gofiber/fiber/v2"

func Running(c *fiber.Ctx) error{
	return c.SendString("api is running")
}