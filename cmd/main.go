package main

import (
	"fmt"
	"github.com/Cheap-NFT-Marketplace/wallet-information/cmd/config"
	"log"
)

func main() {
	serviceConfig, err := config.ReadConfig()
	if err != nil {
		log.Fatal("error trying to start wallet-information service: ", err.Error())
	}

	fmt.Println(serviceConfig)
}
