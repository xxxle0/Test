package controllers

import (
	"github.com/duybkit13/api/services"
	"github.com/gin-gonic/gin"
)

type ScanResultControllerI interface {
	CreateScanResult(c *gin.Context)
	GetScanResultList(c *gin.Context)
	GetScanResultDetail(c *gin.Context)
	UpdateScanResult(c *gin.Context)
	DeleteScanResult(c *gin.Context)
}

type ScanResultController struct {
	scanResultService services.ScanResultServiceI
}

func ScanResultInit(scanResultService services.ScanResultServiceI) ScanResultControllerI {
	return &ScanResultController{
		scanResultService: scanResultService,
	}
}

func (s *ScanResultController) CreateScanResult(c *gin.Context) {

}

func (s *ScanResultController) GetScanResultList(c *gin.Context) {

}

func (s *ScanResultController) GetScanResultDetail(c *gin.Context) {

}

func (s *ScanResultController) UpdateScanResult(c *gin.Context) {

}

func (s *ScanResultController) DeleteScanResult(c *gin.Context) {

}
