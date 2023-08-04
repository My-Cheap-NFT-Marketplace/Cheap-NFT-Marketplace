package config

import (
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	ServiceName      string
	Environment      string
	Port             string
	Url              string
	NftContracts     NftContracts
	AuctionContracts AuctionContracts
}

type NftContracts struct {
	MockERC721 string
}

type AuctionContracts struct {
	Marketplace string
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
	url := viper.GetString("SEPOLIA_URL")

	config.Url = strings.Replace(config.Url, "{SEPOLIA_URL}", url, 1)
	config.Environment = strings.Replace(config.Environment, "{ENVIRONMENT}", env, 1)
	return config, nil
}
