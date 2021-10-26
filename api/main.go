package main

import (
	"log"

	"github.com/duybkit13/api/configurations"
	"github.com/duybkit13/api/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	_, err := configurations.LoadConfig(".")
	if err != nil {
		log.Fatal("Load Config Fail")
	}
	r := gin.Default()
	r.GET("/health-check", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Alive",
		})
	})
	api := r.Group("/api")
	routers.InitResultRouter(api)
	r.Run()
}
