package api

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func Server() {

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	//config.AllowHeaders = []string{"Origin", "Authorization", "Content-Length", "Content-Type"}
	//config.ExposeHeaders = []string{"Content-Length"}
	//config.AllowOrigins = []string{"*"}

	router := gin.Default()
	router.Use(cors.New(config))
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.POST("charge/paystack", GetPaymentLink)
	router.POST("charge/seerbit", GetPaymentLink)

	router.NoRoute(func(c *gin.Context) {
		c.IndentedJSON(404, gin.H{"message": "Selected Service not found"})
	})

	port := os.Getenv("PORT")

	if port == "" {
		port = "8888"
	}
	err := router.Run(fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(err)
		return
	}
}
