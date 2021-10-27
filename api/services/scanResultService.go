package services

import (
	"context"
	"log"

	"github.com/duybkit13/api/dtos"
	"github.com/duybkit13/api/entities"
	"github.com/duybkit13/api/repositories"
)

type ScanResultServiceI interface {
	CreateScanResult(ctx context.Context, createScanResultDto dtos.CreateScanResultDto) error
	GetScanResultDetail(ctx context.Context, getScanResultDetailDto dtos.GetScanResultDetailDto) (dtos.GetScanResultDetailResp, error)
	GetScanResultList(ctx context.Context, getScanResultListDto dtos.GetScanResultListDto) (dtos.GetScanResultListResp, error)
	DeleteScanResult(ctx context.Context, deleteScanResultDto dtos.DeleteScanResultDto) error
	UpdateScanResult(ctx context.Context, updateScanResultDto dtos.UpdateScanResultDto) error
}

type ScanResultService struct {
	scanResultRepository repositories.ScanResultRepositoryI
}

func InitScanResultService(scanResultRepository repositories.ScanResultRepositoryI) ScanResultServiceI {
	return &ScanResultService{
		scanResultRepository: scanResultRepository,
	}
}

func (s *ScanResultService) CreateScanResult(ctx context.Context, createScanResultDto dtos.CreateScanResultDto) error {
	scanResult := entities.Result{
		Status:         createScanResultDto.Status,
		RepositoryName: createScanResultDto.RepositoryName,
		ScanningAt:     createScanResultDto.ScanningAt,
		QueuedAt:       createScanResultDto.QueuedAt,
		FinishedAt:     createScanResultDto.FinishedAt,
	}
	err := s.scanResultRepository.Create(ctx, scanResult)
	return err
}

func (s *ScanResultService) GetScanResultDetail(ctx context.Context, getScanResultDetailDto dtos.GetScanResultDetailDto) (dtos.GetScanResultDetailResp, error) {
	var scanResultResp dtos.ScanResultResp
	scanResult, err := s.GetScanResultDetailById(ctx, getScanResultDetailDto.ID)
	if err != nil {
		return dtos.GetScanResultDetailResp{}, err
	}
	scanResultResp = dtos.ScanResultResp{
		ScanningAt:     scanResult.ScanningAt,
		Status:         scanResult.Status,
		RepositoryName: scanResult.RepositoryName,
		QueuedAt:       scanResult.QueuedAt,
	}
	getScanResultResp := dtos.GetScanResultDetailResp{
		ScanResult: scanResultResp,
	}
	return getScanResultResp, nil
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

func (s *ScanResultService) GetScanResultList(ctx context.Context, getScanResultListDto dtos.GetScanResultListDto) (dtos.GetScanResultListResp, error) {
	condition := map[string]interface{}{}
	count := s.scanResultRepository.Count(ctx, condition)
	scanResults, err := s.scanResultRepository.FindMany(ctx, condition)
	log.Println(scanResults)
	if err != nil {
		return dtos.GetScanResultListResp{}, err
	}
	getScanResultListResp := dtos.GetScanResultListResp{
		Total: count,
	}
	return getScanResultListResp, nil
}

func (s *ScanResultService) DeleteScanResult(ctx context.Context, deleteScanResultDto dtos.DeleteScanResultDto) error {
	condition := map[string]interface{}{
		"id": deleteScanResultDto.ID,
	}
	err := s.scanResultRepository.Delete(ctx, condition)
	return err
}

func (s *ScanResultService) UpdateScanResult(ctx context.Context, updateScanResultDto dtos.UpdateScanResultDto) error {
	condition := map[string]interface{}{}
	newPayload := entities.Result{}
	err := s.scanResultRepository.Update(ctx, condition, newPayload)
	return err
}
