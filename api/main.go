package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	database "github.com/rupadas/raven/config"
	"github.com/rupadas/raven/routes"
)

func setUpRoutes(app *fiber.App) {
	app.Post("/apps", routes.AddApp)
	app.Post("/channels", routes.AddChannel)
	app.Post("/customers", routes.AddCustomer)
	app.Post("/events", routes.AddEvent)
	app.Post("events/:id/channels", routes.AddEventChannels)
	// app.Get("/book/:id", routes.GetBook)
	// app.Post("/book", routes.AddBook)
	// app.Put("/book/:id", routes.Update)
	// app.Delete("/book/:id", routes.Delete)
}

func main() {
	_, err := database.ConnectDb()
	if err != nil {
		log.Fatal(err)
	}
	app := fiber.New()
	setUpRoutes(app)
	app.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))
	log.Fatal(app.Listen(":" + os.Getenv("APP_PORT")))
}
