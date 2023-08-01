package config

import (
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	ServiceName string
	Environment string
	Port        string
	Database    Database
}

type Database struct {
	Marketplace Platform
}

type Platform struct {
	DriverName     string
	DataSourceName string
}

func ReadConfig() (Config, error) {
	viper.SetConfigFile("./cmd/config/config.json")
	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return Config{}, err
	}

	viper.AutomaticEnv()
	env := viper.GetString("ENVIRONMENT")
	marketplaceDbUser := viper.GetString("MARKETPLACE_USER")
	marketplaceDbPassword := viper.GetString("MARKETPLACE_PASSWORD")
	marketplaceHost := viper.GetString("MARKETPLACE_HOST")

	config.Database.Marketplace.DataSourceName = strings.Replace(
		config.Database.Marketplace.DataSourceName, "{MARKETPLACE_USER}", marketplaceDbUser, 1)
	config.Database.Marketplace.DataSourceName = strings.Replace(
		config.Database.Marketplace.DataSourceName, "{MARKETPLACE_PASSWORD}", marketplaceDbPassword, 1)
	config.Database.Marketplace.DataSourceName = strings.Replace(
		config.Database.Marketplace.DataSourceName, "{MARKETPLACE_HOST}", marketplaceHost, 1)
	config.Environment = strings.Replace(config.Environment, "{ENVIRONMENT}", env, 1)

	return config, nil
}
