package services

import (
	"context"

	"github.com/duybkit13/api/adapters"
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
	scanResult, err := adapters.ScanResultAdapter(createScanResultDto)
	if err != nil {
		return err
	}
	_, err = s.scanResultRepository.Create(ctx, scanResult)
	return err
}

func (s *ScanResultService) GetScanResultDetail(ctx context.Context, getScanResultDetailDto dtos.GetScanResultDetailDto) (dtos.GetScanResultDetailResp, error) {
	scanResult, err := s.GetScanResultDetailById(ctx, getScanResultDetailDto.ID)
	if err != nil {
		return dtos.GetScanResultDetailResp{}, err
	}
	scanResultResp, err := adapters.ScanResultRespAdapter(scanResult)
	if err != nil {
		return dtos.GetScanResultDetailResp{}, err
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
	scanResults, err := s.scanResultRepository.FindMany(ctx, condition, getScanResultListDto.Limit, getScanResultListDto.Offset)
	if err != nil {
		return dtos.GetScanResultListResp{}, err
	}
	scanResultListResp := make([]dtos.ScanResultResp, len(scanResults))
	for i, scanResult := range scanResults {
		scanResultResp, err := adapters.ScanResultRespAdapter(scanResult)
		if err != nil {
			return dtos.GetScanResultListResp{}, err
		}
		scanResultListResp[i] = scanResultResp
	}
	getScanResultListResp := dtos.GetScanResultListResp{
		Total:       count,
		ScanResults: scanResultListResp,
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
