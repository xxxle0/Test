package main

import (
	"log"

	"github.com/duybkit13/api/configurations"
	"github.com/duybkit13/api/documentations"
	"github.com/duybkit13/api/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	config, err := configurations.LoadConfig(".")
	if err != nil {
		log.Fatal("Load Config Fail")
	}
	r := gin.Default()
	documentations.InitSwagger(r)
	routers.InitResultRouter(r)
	r.Run(config.ServerAddress)
}
