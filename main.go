package main

import (
	"Payment-Service/api"
	"Payment-Service/config"
)

func init() {
	config.InitDb()
	config.InitServices()
}

func main() {
	api.Server()
}
