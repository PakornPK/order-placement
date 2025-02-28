package route

import (
	"github.com/PakornPK/order-placement/handler"
	"github.com/PakornPK/order-placement/logs"
	"github.com/PakornPK/order-placement/service"
	"github.com/gofiber/fiber/v2"
)

func New(app *fiber.App, logger logs.Logger) {
	serv := service.NewOrderService()
	hand := handler.NewOrderHandler(serv)
	app.Post("/order", hand.PlaceOrderHandler)
}
