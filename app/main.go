package main

import (
	"log"
	"pembayaran/configs"
	"pembayaran/routes"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`app/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	configs.InitDB()
	e := routes.New()
	e.Start(":8000")
}
