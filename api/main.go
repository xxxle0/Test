package main

import (
	"log"
	"time"

	"github.com/duybkit13/api/configurations"
	"github.com/duybkit13/api/databases"
	"github.com/duybkit13/api/documentations"
	"github.com/duybkit13/api/routers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config, err := configurations.LoadConfig(".")
	if err != nil {
		log.Println("Load Config Fail")
	}
	db, err := databases.InitPostgresQLClient(config)
	if err != nil {
		log.Fatal("Connect Database Fail")
	}
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))
	documentations.InitSwagger(r)
	routers.InitResultRouter(r, db)
	r.Run(config.ServerAddress)
}
