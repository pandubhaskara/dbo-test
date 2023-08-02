package database

import (
	"fmt"
	"net/url"
	"os"

	"dbo-backend/models"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB
var dbLogLevel logger.LogLevel

type Database struct {
	Connection string
	Host       string
	Port       string
	Name       string
	User       string
	Password   string
}

var database Database

func MakeConnection() {
	var db *gorm.DB
	var err error

	database = Database{
		Connection: os.Getenv("DB_CONNECTION"),
		Host:       os.Getenv("DB_HOST"),
		Port:       os.Getenv("DB_PORT"),
		Name:       os.Getenv("DB_DATABASE"),
		User:       os.Getenv("DB_USERNAME"),
		Password:   os.Getenv("DB_PASSWORD"),
	}

	connection := database.Connection
	fmt.Printf("Initializing database connection " + connection)

	switch connection {
	case "postgres":
		db, err = pg()
	case "mysql":
		fallthrough
	default:
		db, err = my()
	}

	if err != nil {
		fmt.Println(err)
		panic("DB Connection Failed.")
	}

	Db = db

	db.AutoMigrate(
		&models.User{},
		&models.AccessToken{},
		&models.RefreshToken{},
		&models.Order{},
		&models.Product{},
		&models.Supplier{},
	)
}

func my() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=%s", database.User, database.Password, database.Host, database.Port, database.Name, url.QueryEscape(os.Getenv("APP_TIMEZONE")))
	return gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(dbLogLevel),
	})
}

func pg() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", database.Host, database.User, database.Password, database.Name, database.Port, os.Getenv("APP_TIMEZONE"))
	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:         logger.Default.LogMode(dbLogLevel),
		TranslateError: true,
	})
}
