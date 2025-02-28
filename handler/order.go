package handler

import "github.com/gofiber/fiber/v2"

type OrderHandler interface {
	PlaceOrderHandler(c *fiber.Ctx) error
}

type orderHandler struct {
}

func NewOrderHandler() OrderHandler {
	return orderHandler{}
}

func (h orderHandler) PlaceOrderHandler(c *fiber.Ctx) error {
	
	return c.Status(fiber.StatusOK).JSON(map[string]string{"status": "OK"})
}
