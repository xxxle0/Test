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

// CreateScanResult godoc
// @Summary Create Scan Result
// @Description Create Scan Result With ScanResult Payload
// @Accept  json
// @Produce  json
// @Header 200 {string} Token "qwerty"
// @Router /v1/scan-results [post]
func (s *ScanResultController) CreateScanResult(c *gin.Context) {

}

// GetScanResultList godoc
// @Summary Get Scan Result List
// @Description Get Scan Result List with Offset & Payload
// @Accept  json
// @Produce  json
// @Header 200 {string} Token "qwerty"
// @Router /v1/scan-results [get]
func (s *ScanResultController) GetScanResultList(c *gin.Context) {

}

// GetScanResultDetail godoc
// @Summary Get Scan Result Detail
// @Description Get Scan Result Detail By ID
// @Accept  json
// @Produce  json
// @Param id path int true "Scan Result ID"
// @Header 200 {string} Token "qwerty"
// @Router /v1/scan-results/:id [get]
func (s *ScanResultController) GetScanResultDetail(c *gin.Context) {

}

// UpdateScanResult godoc
// @Summary Update Scan Result
// @Description Update Scan Result
// @Accept  json
// @Produce  json
// @Param id path int true "Scan Result ID"
// @Header 200 {string} Token "qwerty"
// @Router /v1/scan-results [patch]
func (s *ScanResultController) UpdateScanResult(c *gin.Context) {

}

// DeleteScanResult godoc
// @Summary Delete Scan Result
// @Description Delete Scan Result
// @Accept  json
// @Produce  json
// @Param id path int true "Scan Result ID"
// @Header 200 {string} Token "qwerty"
// @Router /v1/scan-results/:id [delete]
func (s *ScanResultController) DeleteScanResult(c *gin.Context) {

}
