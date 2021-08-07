package main

import (
	"pembayaran/configs"
	"pembayaran/routes"
)

func main() {
	configs.InitDB()
	e := routes.New()
	e.Start(":8000")
}
