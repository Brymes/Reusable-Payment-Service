package main

import (
	"Payment-Service/api"
	"Payment-Service/services"
)

func init() {
	//config.InitDb()
	services.InitServices()
}

func main() {
	api.Server()
}
