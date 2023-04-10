package controllers

import "github.com/gofiber/fiber/v2"

func Hello(context *fiber.Ctx) error {
	response := "Hello World!"
	return context.SendString(response)
}
