package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	database "github.com/rupadas/notify/config"
	"github.com/rupadas/notify/routes"
)

func main() {
	_, err := database.ConnectDb()
	if err != nil {
		log.Fatal(err)
	}
	app := fiber.New()
	routes.SetUpRoutes(app)
	app.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))
	log.Fatal(app.Listen(":" + os.Getenv("APP_PORT")))
}
