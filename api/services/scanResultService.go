package services

import (
	"context"
	"log"

	"github.com/duybkit13/api/dtos"
	"github.com/duybkit13/api/entities"
	"github.com/duybkit13/api/repositories"
)

type ScanResultServiceI interface {
	CreateScanResult(ctx context.Context, createScanResultDto dtos.CreateScanResultDto)
	GetScanResultDetail(ctx context.Context, getScanResultDetailDto dtos.GetScanResultDetailDto)
	GetScanResultList(ctx context.Context, getScanResultListDto dtos.GetScanResultListDto)
	DeleteScanResult(ctx context.Context, deleteScanResultDto dtos.DeleteScanResultDto)
	UpdateScanResult(ctx context.Context, updateScanResultDto dtos.UpdateScanResultDto)
}

type ScanResultService struct {
	scanResultRepository repositories.ScanResultRepositoryI
}

func InitScanResultService(scanResultRepository repositories.ScanResultRepositoryI) ScanResultServiceI {
	return &ScanResultService{
		scanResultRepository: scanResultRepository,
	}
}

func (s *ScanResultService) CreateScanResult(ctx context.Context, createScanResultDto dtos.CreateScanResultDto) {
	scanResult := entities.Result{
		Status:         createScanResultDto.Status,
		RepositoryName: createScanResultDto.RepositoryName,
		ScanningAt:     createScanResultDto.ScanningAt,
		QueuedAt:       createScanResultDto.QueuedAt,
		FinishedAt:     createScanResultDto.FinishedAt,
	}
	s.scanResultRepository.Create(ctx, scanResult)
}

func (s *ScanResultService) GetScanResultDetail(ctx context.Context, getScanResultDetailDto dtos.GetScanResultDetailDto) {
	scanResult, err := s.GetScanResultDetailById(ctx, getScanResultDetailDto.ID)
	if err != nil {

	}
	log.Println(scanResult)
}

func (s *ScanResultService) GetScanResultDetailById(ctx context.Context, id int) (entities.Result, error) {
	condition := map[string]interface{}{
		"id": id,
	}
	scanResult, err := s.scanResultRepository.FindOne(ctx, condition)
	if err != nil {
		return scanResult, err
	}
	return scanResult, nil
}

func (s *ScanResultService) GetScanResultList(ctx context.Context, getScanResultListDto dtos.GetScanResultListDto) {
	condition := map[string]interface{}{}
	count := s.scanResultRepository.Count(ctx, condition)
	scanResults, err := s.scanResultRepository.FindMany(ctx, condition)
	if err != nil {

	}
	log.Println(scanResults, count)
}

func (s *ScanResultService) DeleteScanResult(ctx context.Context, deleteScanResultDto dtos.DeleteScanResultDto) {
	condition := map[string]interface{}{
		"id": deleteScanResultDto.ID,
	}
	err := s.scanResultRepository.Delete(ctx, condition)
	if err != nil {

	}
}

func (s *ScanResultService) UpdateScanResult(ctx context.Context, updateScanResultDto dtos.UpdateScanResultDto) {
	condition := map[string]interface{}{}
	newPayload := entities.Result{}
	err := s.scanResultRepository.Update(ctx, condition, newPayload)
	if err != nil {

	}
}
