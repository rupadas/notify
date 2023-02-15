package handler

import (
	"crypto/rand"
	"encoding/hex"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	database "github.com/rupadas/notify/config"
	"github.com/rupadas/notify/models"
)

func AddApp(c *fiber.Ctx) error {
	account := new(models.App)
	if err := c.BodyParser(account); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	userId, ok := c.Locals("userId").(uint)
	if !ok {
		log.Println(ok)
	}
	account.UserId = userId

	// Generate a 32-byte key
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}

	// Generate a 64-byte token
	token := make([]byte, 64)
	_, err = rand.Read(token)
	if err != nil {
		panic(err)
	}
	account.AccessKey = hex.EncodeToString(key)
	account.AccessToken = hex.EncodeToString((token))
	database.DBConn.Create(&account)
	return c.Status(http.StatusOK).JSON(account)
}

func GetApp(AcessKey, AccessToken string) (*models.App, error) {
	app := &models.App{}
	// query the database to retrieve the apps where the access key and access token match
	err := database.DBConn.Where("access_key = ? AND access_token = ?", AcessKey, AccessToken).First(app).Error
	if err != nil {
		return nil, err
	}
	return app, nil
}
