package main

import (
	"fmt"

	"github.com/PakornPK/order-placement/config"
	"github.com/PakornPK/order-placement/database"
	"github.com/PakornPK/order-placement/logs"
	"github.com/PakornPK/order-placement/route"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func main() {
	conf := config.LoadConfig()
	logger, sync := logs.NewLogger(conf.App)
	defer sync()
	db := database.NewDatabase(conf.Db)

	if db != nil {
		logger.Info("init database success")
	}
	app := fiber.New()
	route.New(app, *db, logger)
	log.Fatal(app.Listen(fmt.Sprintf(":%d", conf.App.Port)))
}
