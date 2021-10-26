package main

import (
	"log"

	"github.com/duybkit13/api/configurations"
	"github.com/duybkit13/api/databases"
	"github.com/duybkit13/api/documentations"
	"github.com/duybkit13/api/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	config, err := configurations.LoadConfig(".")
	if err != nil {
		log.Fatal("Load Config Fail")
	}
	db, err := databases.InitPostgresQLClient(config)
	if err != nil {
		log.Fatal("Connect Database Fail")
	}
	r := gin.Default()
	documentations.InitSwagger(r)
	routers.InitResultRouter(r, db)
	r.Run(config.ServerAddress)
}
