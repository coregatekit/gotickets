package database

import (
	"fmt"
	"log"

	"github.com/coregate/tickets-app/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConnection struct {
	DB *gorm.DB
}

type Config struct {
	Host      string
	Port      int32
	User      string
	Password  string
	Database  string
	EnableTLS bool
}

func NewPostgres(configs *configs.Configs) (*DatabaseConnection, error) {
	connection, err := openConnectionPostgres(Config{
		Host:     configs.Database.Host,
		Port:     configs.Database.Port,
		User:     configs.Database.User,
		Password: configs.Database.Password,
		Database: configs.Database.Name,
	})
	if err != nil {
		return nil, err
	}

	return &DatabaseConnection{DB: connection}, nil
}

func openConnectionPostgres(config Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Database)
	fmt.Println(dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	log.Println("Database connection established")

	return db, nil
}
