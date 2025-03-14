package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/coregate/tickets-app/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,  // Slow SQL threshold
			LogLevel:                  logger.Error, // Log level
			IgnoreRecordNotFoundError: true,         // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,         // Don't include params in the SQL log
			Colorful:                  true,         // Disable color
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	log.Println("Database connection established")

	return db, nil
}
