package controllers

import "github.com/duybkit13/api/services"

type ScanResultControllerI interface {
}

type ScanResultController struct {
	scanResultService services.ScanResultServiceI
}

func ScanResultInit(scanResultService services.ScanResultServiceI) ScanResultControllerI {
	return &ScanResultController{
		scanResultService: scanResultService,
	}
}

func (s *ScanResultController) CreateScanResult() {

}

func (s *ScanResultController) GetScanResultList() {

}

func (s *ScanResultController) GetScanResultDetail() {

}

func (s *ScanResultController) UpdateScanResult() {

}

func (s *ScanResultController) DeleteScanResult() {

}
