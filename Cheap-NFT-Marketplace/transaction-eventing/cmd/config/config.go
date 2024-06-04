package config

import (
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	ServiceName   string
	Environment   string
	Port          string
	Database      Database
	Nats          Nats
	ElasticSearch ElasticSearch
	Topics        []string
}

type Database struct {
	Marketplace Platform
}

type Platform struct {
	DriverName     string
	DataSourceName string
}

type Nats struct {
	Host string
}

type ElasticSearch struct {
	Host string
}

func ReadConfig() (Config, error) {
	viper.SetConfigFile("./cmd/config/config.json")
	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}

	viper.AutomaticEnv()

	env := viper.GetString("ENVIRONMENT")
	marketplaceDbUser := viper.GetString("MARKETPLACE_USER")
	marketplaceDbPassword := viper.GetString("MARKETPLACE_PASSWORD")
	marketplaceHost := viper.GetString("MARKETPLACE_HOST")
	natsHost := viper.GetString("NATS_HOST")
	elasticHost := viper.GetString("ELASTIC_HOST")

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return Config{}, err
	}

	config.ServiceName = strings.Replace(config.ServiceName, "${ENVIRONMENT}", env, 1)
	config.Database.Marketplace.DataSourceName = strings.Replace(config.Database.Marketplace.DataSourceName, "{MARKETPLACE_USER}", marketplaceDbUser, 1)
	config.Database.Marketplace.DataSourceName = strings.Replace(config.Database.Marketplace.DataSourceName, "{MARKETPLACE_PASSWORD}", marketplaceDbPassword, 1)
	config.Database.Marketplace.DataSourceName = strings.Replace(config.Database.Marketplace.DataSourceName, "{MARKETPLACE_HOST}", marketplaceHost, 1)
	config.Nats.Host = strings.Replace(config.Nats.Host, "${NATS_HOST}", natsHost, 1)
	config.ElasticSearch.Host = strings.Replace(config.ElasticSearch.Host, "${ELASTIC_HOST}", elasticHost, 1)
	return config, nil
}
