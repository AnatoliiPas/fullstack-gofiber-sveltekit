package controllers

import "github.com/gofiber/fiber/v2"

func Index(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"title":       "Главная",
		"description": "главная страница",
	})
}
