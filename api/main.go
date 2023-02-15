package main

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/golang-jwt/jwt"
	database "github.com/rupadas/raven/config"
	"github.com/rupadas/raven/models"
	"github.com/rupadas/raven/routes"
)

func getApp(AcessKey, AccessToken string) (*models.App, error) {
	app := &models.App{}
	// query the database to retrieve the apps where the access key and access token match
	err := database.DBConn.Where("access_key = ? AND access_token = ?", AcessKey, AccessToken).First(app).Error
	if err != nil {
		return nil, err
	}
	return app, nil
}

func authenticationMiddleware(c *fiber.Ctx) error {
	AcessKey := c.Get("AcessKey")
	AccessToken := c.Get("AccessToken")
	app, error := getApp(AcessKey, AccessToken)
	log.Println(error)
	c.Locals("Environment", app.Environment)
	c.Locals("appId", app.ID)
	err := c.Next()
	return err
}

type JWTClaim struct {
	UserId uint   `json:"user_id"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

func jwtMiddleware(c *fiber.Ctx) error {
	signedToken := c.Get("x-token")
	if signedToken == "" {
		err := errors.New("couldn't find x-token")
		return err
	}
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte("supersecretkey"), nil
		},
	)
	if err != nil {
		return err
	}
	claims, ok := token.Claims.(*JWTClaim)
	c.Locals("userId", claims.UserId)
	if !ok {
		err = errors.New("couldn't parse claims")
		return err
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return err
	}
	error := c.Next()
	return error
}

func setUpRoutes(app *fiber.App) {
	app.Post("/signup", routes.SignUp)
	app.Post("/login", routes.SignIn)
	app.Post("/apps", jwtMiddleware, routes.AddApp)
	app.Post("/channels", authenticationMiddleware, routes.AddChannel)
	app.Post("/customers", routes.AddCustomer)
	app.Post("/events", authenticationMiddleware, routes.AddEvent)
	app.Post("events/:id/channels", authenticationMiddleware, routes.AddEventChannels)
	app.Post("/providers", authenticationMiddleware, routes.AddProvider)
	app.Post("/channels/:id/providers", authenticationMiddleware, routes.AddChannelProvider)
	app.Post("/events/sendEmail", authenticationMiddleware, routes.SendEmail)
	app.Post("/events/sendSms", authenticationMiddleware, routes.SendSms)
	app.Post("/providers/:id/setting", authenticationMiddleware, routes.AddProviderSetting)
	app.Get("events/:id/channels", authenticationMiddleware, routes.FetchEventChannles)
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
