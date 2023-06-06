package main

import (
	"fmt"
	"log"

	conf "github.com/acework2u/air-iot-app-api-service/config"
	"github.com/gin-gonic/gin"
)

func main() {

	configs, err := conf.LoadCongig(".")

	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":  "OK API",
			"data env": fmt.Sprintf("Port :", configs.Port),
		})
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":3000") // listen and serve on 0.0.0.0:8080

}
