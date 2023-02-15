package middleware

import (
	"errors"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/rupadas/notify/handler"
)

type JWTClaim struct {
	UserId uint   `json:"user_id"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

func JwtMiddleware(c *fiber.Ctx) error {
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

func AuthenticationMiddleware(c *fiber.Ctx) error {
	AcessKey := c.Get("AcessKey")
	AccessToken := c.Get("AccessToken")
	app, error := handler.GetApp(AcessKey, AccessToken)
	log.Println(error)
	c.Locals("Environment", app.Environment)
	c.Locals("appId", app.ID)
	err := c.Next()
	return err
}
