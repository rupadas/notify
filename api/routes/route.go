package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rupadas/notify/handler"
	"github.com/rupadas/notify/middleware"
)

func SetUpRoutes(app *fiber.App) {
	app.Post("/signup", handler.SignUp)
	app.Post("/login", handler.SignIn)
	app.Post("/apps", middleware.JwtMiddleware, handler.AddApp)
	app.Post("/channels", middleware.AuthenticationMiddleware, handler.AddChannel)
	app.Post("/events", middleware.AuthenticationMiddleware, handler.AddEvent)
	app.Post("events/:id/channels", middleware.AuthenticationMiddleware, handler.AddEventChannels)
	app.Post("/providers", middleware.AuthenticationMiddleware, handler.AddProvider)
	app.Post("/channels/:id/providers", middleware.AuthenticationMiddleware, handler.AddChannelProviders)
	app.Post("/providers/:id/settings", middleware.AuthenticationMiddleware, handler.AddProviderSettings)
	app.Get("events/:id/channels", middleware.AuthenticationMiddleware, handler.FetchEventChannles)
	app.Post("/events/sendEmail", middleware.AuthenticationMiddleware, handler.SendEmail)
	app.Post("/events/sendSms", middleware.AuthenticationMiddleware, handler.SendSms)
}
