package services

import (
	"log"

	"github.com/duybkit13/api/dtos"
	"github.com/duybkit13/api/entities"
	"github.com/duybkit13/api/repositories"
)

type ScanResultServiceI interface {
	CreateScanResult(createScanResultDto dtos.CreateScanResultDto)
	GetScanResultDetail(getScanResultDetailDto dtos.GetScanResultDetailDto)
	GetScanResultList(getScanResultListDto dtos.GetScanResultListDto)
	DeleteScanResult(deleteScanResultDto dtos.DeleteScanResultDto)
	UpdateScanResult(updateScanResultDto dtos.UpdateScanResultDto)
}

type ScanResultService struct {
	scanResultRepository repositories.ScanResultRepositoryI
}

func InitScanResultService(scanResultRepository repositories.ScanResultRepositoryI) ScanResultServiceI {
	return &ScanResultService{
		scanResultRepository: scanResultRepository,
	}
}

func (s *ScanResultService) CreateScanResult(createScanResultDto dtos.CreateScanResultDto) {
	scanResult := entities.Result{
		Status:         createScanResultDto.Status,
		RepositoryName: createScanResultDto.RepositoryName,
		ScanningAt:     createScanResultDto.ScanningAt,
		QueuedAt:       createScanResultDto.QueuedAt,
		FinishedAt:     createScanResultDto.FinishedAt,
	}
	s.scanResultRepository.Create(scanResult)
}

func (s *ScanResultService) GetScanResultDetail(getScanResultDetailDto dtos.GetScanResultDetailDto) {
	condition := map[string]interface{}{
		"id": getScanResultDetailDto.ID,
	}
	scanResult, err := s.scanResultRepository.FindOne(condition)
	if err != nil {

	}
	log.Println(scanResult)
}

func (s *ScanResultService) GetScanResultList(getScanResultListDto dtos.GetScanResultListDto) {

}

func (s *ScanResultService) DeleteScanResult(deleteScanResultDto dtos.DeleteScanResultDto) {
	condition := map[string]interface{}{
		"id": deleteScanResultDto.ID,
	}
	err := s.scanResultRepository.Delete(condition)
	if err != nil {

	}
}

func (s *ScanResultService) UpdateScanResult(updateScanResultDto dtos.UpdateScanResultDto) {
	condition := map[string]interface{}{}
	newPayload := entities.Result{}
	err := s.scanResultRepository.Update(condition, newPayload)
	if err != nil {

	}
}
