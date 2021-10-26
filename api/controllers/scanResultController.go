package controllers

import (
	"github.com/duybkit13/api/dtos"
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

// @Tags Scan Results
// CreateScanResult godoc
// @Summary Create Scan Result
// @Description Create Scan Result With ScanResult Payload
// @Accept  json
// @Produce  json
// @Header 200 {string} Token "qwerty"
// @Router /v1/scan-results [post]
func (s *ScanResultController) CreateScanResult(c *gin.Context) {
	var createScanResultDto dtos.CreateScanResultDto
	err := c.ShouldBindJSON(&createScanResultDto)
	if err != nil {

	}
	s.scanResultService.CreateScanResult(createScanResultDto)
}

// @Tags Scan Results
// GetScanResultList godoc
// @Summary Get Scan Result List
// @Description Get Scan Result List with Offset & Payload
// @Accept  json
// @Produce  json
// @Header 200 {string} Token "qwerty"
// @Router /v1/scan-results [get]
func (s *ScanResultController) GetScanResultList(c *gin.Context) {
	var getScanResultListDto dtos.GetScanResultListDto
	err := c.ShouldBindQuery(getScanResultListDto)
	if err != nil {

	}
	s.scanResultService.GetScanResultList(getScanResultListDto)
}

// @Tags Scan Results
// GetScanResultDetail godoc
// @Summary Get Scan Result Detail
// @Description Get Scan Result Detail By ID
// @Accept  json
// @Produce  json
// @Param id path int true "Scan Result ID"
// @Header 200 {string} Token "qwerty"
// @Router /v1/scan-results/:id [get]
func (s *ScanResultController) GetScanResultDetail(c *gin.Context) {
	var getScanResultDetailDto dtos.GetScanResultDetailDto
	err := c.ShouldBindQuery(getScanResultDetailDto)
	if err != nil {

	}
	s.scanResultService.GetScanResultDetail(getScanResultDetailDto)
}

// @Tags Scan Results
// UpdateScanResult godoc
// @Summary Update Scan Result
// @Description Update Scan Result
// @Accept  json
// @Produce  json
// @Param id path int true "Scan Result ID"
// @Header 200 {string} Token "qwerty"
// @Router /v1/scan-results [patch]
func (s *ScanResultController) UpdateScanResult(c *gin.Context) {
	var updateScanResultDto dtos.UpdateScanResultDto
	err := c.ShouldBindJSON(updateScanResultDto)
	if err != nil {

	}
	s.scanResultService.UpdateScanResult(updateScanResultDto)
}

// @Tags Scan Results
// DeleteScanResult godoc
// @Summary Delete Scan Result
// @Description Delete Scan Result
// @Accept  json
// @Produce  json
// @Param id path int true "Scan Result ID"
// @Header 200 {string} Token "qwerty"
// @Router /v1/scan-results/:id [delete]
func (s *ScanResultController) DeleteScanResult(c *gin.Context) {
	var deleteScanResultDto dtos.DeleteScanResultDto
	err := c.ShouldBindUri(deleteScanResultDto)
	if err != nil {

	}
	s.scanResultService.DeleteScanResult(deleteScanResultDto)
}
