package configs

import "github.com/spf13/viper"

type (
	Configs struct {
		App      *App
		Database *Database
	}

	App struct {
		Name string
		Env  string
		Port int32
	}

	Database struct {
		Host     string
		Port     int32
		User     string
		Password string
		Name     string
	}
)

func NewConfigs() *Configs {
	viper.AutomaticEnv()

	app := &App{
		Name: viper.GetString("APP_NAME"),
		Env:  viper.GetString("APP_ENV"),
		Port: viper.GetInt32("APP_PORT"),
	}

	db := &Database{
		Host:     viper.GetString("DB_HOST"),
		Port:     viper.GetInt32("DB_PORT"),
		User:     viper.GetString("DB_USER"),
		Password: viper.GetString("DB_PASSWORD"),
		Name:     viper.GetString("DB_NAME"),
	}

	return &Configs{
		App:      app,
		Database: db,
	}
}
