package route

import (
	"github.com/PakornPK/order-placement/database"
	"github.com/PakornPK/order-placement/handler"
	"github.com/PakornPK/order-placement/logs"
	"github.com/gofiber/fiber/v2"
)

func New(app *fiber.App, db database.Database, logger logs.Logger) {
	hand := handler.NewOrderHandler()
	app.Post("/order", hand.PlaceOrderHandler)
}
