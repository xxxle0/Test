package routers

import (
	"github.com/duybkit13/api/controllers"
	"github.com/duybkit13/api/repositories"
	"github.com/duybkit13/api/services"
	"github.com/gin-gonic/gin"
)

func InitResultRouter(router *gin.RouterGroup) {
	v1 := router.Group("/v1")
	scanResultRouter := v1.Group("/scan-results")
	scanResultRepository := repositories.InitScanResultRepository()
	scanResultService := services.InitScanResultService(scanResultRepository)
	scanResultController := controllers.ScanResultInit(scanResultService)
	{
		scanResultRouter.POST("", scanResultController.CreateScanResult)
		scanResultRouter.GET("", scanResultController.GetScanResultList)
		scanResultRouter.GET("/:id", scanResultController.GetScanResultDetail)
		scanResultRouter.PATCH("/:id", scanResultController.UpdateScanResult)
		scanResultRouter.DELETE("", scanResultController.DeleteScanResult)
	}
}
