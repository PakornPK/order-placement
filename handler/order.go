package handler

import (
	dto "github.com/PakornPK/order-placement/Dto"
	"github.com/PakornPK/order-placement/service"
	"github.com/gofiber/fiber/v2"
)

type OrderHandler interface {
	PlaceOrderHandler(c *fiber.Ctx) error
}

type orderHandler struct {
	serv service.OrderService
}

func NewOrderHandler(serv service.OrderService) OrderHandler {
	return orderHandler{serv: serv}
}

func (h orderHandler) PlaceOrderHandler(c *fiber.Ctx) error {
	req := make([]dto.InputOrder, 0)
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(nil)
	}

	res, err := h.serv.PlaceOrder(req)
	if err != nil {
		return c.Status(fiber.StatusConflict).JSON(nil)
	}
	return c.Status(fiber.StatusOK).JSON(res)
}
