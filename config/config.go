package config

import "github.com/vrischmann/envconfig"

var appconfig Config

type Config struct {
	Database struct {
		URL     string `envconfig:"default=postgres://postgres:mysecretpassword@localhost:5432/postgres?sslmode=disable, optional"`
		Version uint   `envconfig:"default=1"`
	}
	Port       string `envconfig:"default=8080"`
	APPVersion string `envconfig:"default=v1"`
}

func InitConfig() error {
	appconfig = Config{}
	err := envconfig.Init(&appconfig)

	if err != nil {
		return err
	}

	return nil
}

func GetConfig() Config {
	return appconfig
}
