package controllers

import (
	"net/http"

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
// @Router /v1/scan-results [post]
func (s *ScanResultController) CreateScanResult(c *gin.Context) {
	var createScanResultDto dtos.CreateScanResultDto
	err := c.ShouldBindJSON(&createScanResultDto)
	if err != nil {
		response := dtos.Response{
			StatusCode: http.StatusBadRequest,
			ErrorMsg:   err.Error(),
		}
		dtos.BadRequest(c, response)
		return
	}
	err = s.scanResultService.CreateScanResult(c, createScanResultDto)
	if err != nil {
		response := dtos.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Create Scan Result Fail",
		}
		dtos.ErrorResponse(c, response)
		return
	}
	response := dtos.Response{
		StatusCode: http.StatusOK,
		Message:    "Create Scan Result Success",
	}
	dtos.HttpResponse(c, response)
}

// @Tags Scan Results
// GetScanResultList godoc
// @Summary Get Scan Result List
// @Description Get Scan Result List with Offset & Payload
// @Accept  json
// @Produce  json
// @Router /v1/scan-results [get]
func (s *ScanResultController) GetScanResultList(c *gin.Context) {
	var getScanResultListDto dtos.GetScanResultListDto
	err := c.ShouldBindQuery(&getScanResultListDto)
	if err != nil {
		response := dtos.Response{
			StatusCode: http.StatusBadRequest,
			ErrorMsg:   err.Error(),
		}
		dtos.BadRequest(c, response)
		return
	}
	scanResultListResp, err := s.scanResultService.GetScanResultList(c, getScanResultListDto)
	if err != nil {
		response := dtos.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Get Scan Result List Fail",
		}
		dtos.ErrorResponse(c, response)
		return
	}
	response := dtos.Response{
		StatusCode: http.StatusOK,
		Data:       scanResultListResp,
	}
	dtos.HttpResponse(c, response)
}

// @Tags Scan Results
// GetScanResultDetail godoc
// @Summary Get Scan Result Detail
// @Description Get Scan Result Detail By ID
// @Accept  json
// @Produce  json
// @Param id path int true "Scan Result ID"
// @Router /v1/scan-results/:id [get]
func (s *ScanResultController) GetScanResultDetail(c *gin.Context) {
	var getScanResultDetailDto dtos.GetScanResultDetailDto
	err := c.ShouldBindUri(&getScanResultDetailDto)
	if err != nil {
		response := dtos.Response{
			StatusCode: http.StatusBadRequest,
			ErrorMsg:   err.Error(),
		}
		dtos.BadRequest(c, response)
		return
	}
	scanResultDetailResp, err := s.scanResultService.GetScanResultDetail(c, getScanResultDetailDto)
	if err != nil {
		response := dtos.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Get Scan Result Detail Fail",
		}
		dtos.ErrorResponse(c, response)
		return
	}
	response := dtos.Response{
		StatusCode: http.StatusOK,
		Data:       scanResultDetailResp,
	}
	dtos.HttpResponse(c, response)
}

// @Tags Scan Results
// UpdateScanResult godoc
// @Summary Update Scan Result
// @Description Update Scan Result
// @Accept  json
// @Produce  json
// @Param id path int true "Scan Result ID"
// @Router /v1/scan-results [patch]
func (s *ScanResultController) UpdateScanResult(c *gin.Context) {
	var updateScanResultDto dtos.UpdateScanResultDto
	err := c.ShouldBindJSON(&updateScanResultDto)
	if err != nil {
		response := dtos.Response{
			StatusCode: http.StatusBadRequest,
			ErrorMsg:   err.Error(),
		}
		dtos.BadRequest(c, response)
		return
	}
	err = s.scanResultService.UpdateScanResult(c, updateScanResultDto)
	if err != nil {
		response := dtos.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Update Scan Result Fail",
		}
		dtos.ErrorResponse(c, response)
		return
	}
	response := dtos.Response{
		StatusCode: http.StatusOK,
		Message:    "Update Scan Result Success",
	}
	dtos.HttpResponse(c, response)
}

// @Tags Scan Results
// DeleteScanResult godoc
// @Summary Delete Scan Result
// @Description Delete Scan Result
// @Accept  json
// @Produce  json
// @Param id path int true "Scan Result ID"
// @Router /v1/scan-results/:id [delete]
func (s *ScanResultController) DeleteScanResult(c *gin.Context) {
	var deleteScanResultDto dtos.DeleteScanResultDto
	err := c.ShouldBindUri(&deleteScanResultDto)
	if err != nil {
		response := dtos.Response{
			StatusCode: http.StatusBadRequest,
			ErrorMsg:   err.Error(),
		}
		dtos.BadRequest(c, response)
		return
	}
	err = s.scanResultService.DeleteScanResult(c, deleteScanResultDto)
	if err != nil {
		response := dtos.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Delete Scan Result Fail",
		}
		dtos.ErrorResponse(c, response)
		return
	}
	response := dtos.Response{
		StatusCode: http.StatusOK,
		Message:    "Delete Scan Result Success",
	}
	dtos.HttpResponse(c, response)
}
