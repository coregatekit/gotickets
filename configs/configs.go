package configs

import "github.com/spf13/viper"

type (
	Configs struct {
		App      *App
		Database *Database
		Argon    *ArgonParams
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

	ArgonParams struct {
		Memory      uint32
		Iterations  uint32
		Parallelism uint8
		SaltLength  uint32
		KeyLength   uint32
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

	argon := &ArgonParams{
		Memory:      64 * 1024,
		Iterations:  3,
		Parallelism: 2,
		SaltLength:  16,
		KeyLength:   32,
	}

	return &Configs{
		App:      app,
		Database: db,
		Argon:    argon,
	}
}
