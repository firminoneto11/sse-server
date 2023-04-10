package controllers

import "github.com/gofiber/fiber/v2"

func SSEController(context *fiber.Ctx) error {
	response := "Hello World!"
	return context.SendString(response)
}

func NewEventController(context *fiber.Ctx) error {
	response := "Hello World!"
	return context.SendString(response)
}
