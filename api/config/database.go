package config

import (
	"fmt"
	"log"
	"os"

	"github.com/rupadas/raven/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func ConnectDb() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"))
	log.Println("dsn", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		return nil, err
	}
	log.Println("connected", db)
	db.AutoMigrate(&models.App{}, &models.Channel{}, &models.Customer{}, &models.CustomerMeta{}, &models.Event{}, &models.EventChannel{}, &models.Provider{}, &models.ChannelProvider{}, &models.ProviderSetting{})
	DBConn = db
	return db, nil
}
