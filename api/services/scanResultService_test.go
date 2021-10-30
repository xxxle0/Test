package services

import (
	"context"
	"testing"
	"time"

	"github.com/duybkit13/api/adapters"
	"github.com/duybkit13/api/dtos"
	"github.com/duybkit13/api/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockScanResultRepository struct {
	mock.Mock
}

func (m *MockScanResultRepository) Create(ctx context.Context, scanResult entities.Result) error {
	ret := m.Called(ctx, scanResult)
	return ret.Error(0)
}

func (m *MockScanResultRepository) FindOne(ctx context.Context, condition map[string]interface{}) (entities.Result, error) {
	ret := m.Called(ctx, condition)
	return ret.Get(0).(entities.Result), ret.Error(1)
}

func (m *MockScanResultRepository) FindMany(ctx context.Context, condition map[string]interface{}, limit, offset int) ([]entities.Result, error) {
	ret := m.Called(ctx, condition, limit, offset)
	return ret.Get(0).([]entities.Result), ret.Error(1)
}

func (m *MockScanResultRepository) Delete(ctx context.Context, condition map[string]interface{}) error {
	ret := m.Called(ctx, condition)
	return ret.Error(0)
}

func (m *MockScanResultRepository) Update(ctx context.Context, condition map[string]interface{}, newPayload entities.Result) error {
	ret := m.Called(ctx, condition, newPayload)
	return ret.Error(0)
}

func (m *MockScanResultRepository) Count(ctx context.Context, condition map[string]interface{}) int64 {
	ret := m.Called(ctx, condition)
	return int64(ret.Int(0))
}

func TestCreateScanResult(t *testing.T) {
	currentTime := time.Now()
	createScanResultDto := dtos.CreateScanResultDto{
		Status:         "Queued",
		RepositoryName: "Repo Test 1",
		QueuedAt:       currentTime,
		ScanningAt:     currentTime,
		FinishedAt:     currentTime,
	}
	scanResultDto, _ := adapters.ScanResultAdapter(createScanResultDto)
	expectCreateScanResultResp := dtos.CreateScanResultResp{
		Message: "Create Scan Result Success",
	}
	mockScanResultRepository := new(MockScanResultRepository)
	mockScanResultRepository.On("Create", context.Background(), scanResultDto).Return(nil)
	scanResultService := InitScanResultService(mockScanResultRepository)
	createScanResultResp, _ := scanResultService.CreateScanResult(context.Background(), createScanResultDto)
	assert.Equal(t, expectCreateScanResultResp, createScanResultResp)
	mockScanResultRepository.AssertExpectations(t)
}

func TestGetScanResultDetail(t *testing.T) {
	currentTime := time.Now()
	scanResultDto := dtos.GetScanResultDetailDto{
		ID: 1,
	}
	condition := map[string]interface{}{
		"id": scanResultDto.ID,
	}
	mockScanResult := entities.Result{
		Status:         "Queued",
		RepositoryName: "Repo Test 1",
		QueuedAt:       currentTime,
		ScanningAt:     currentTime,
		FinishedAt:     currentTime,
	}
	mockScanResultResp, _ := adapters.ScanResultRespAdapter(mockScanResult)
	expectScanResultResp := dtos.GetScanResultDetailResp{
		ScanResult: mockScanResultResp,
	}
	mockScanResultRepository := new(MockScanResultRepository)
	mockScanResultRepository.On("FindOne", context.Background(), condition).Return(mockScanResult, nil)
	scanResultService := InitScanResultService(mockScanResultRepository)
	scanResultResp, _ := scanResultService.GetScanResultDetail(context.Background(), scanResultDto)
	assert.Equal(t, expectScanResultResp, scanResultResp)
	mockScanResultRepository.AssertExpectations(t)
}

func TestFindManyScanResult(t *testing.T) {
	currentTime := time.Now()
	getScanResultListDto := dtos.GetScanResultListDto{
		Limit:  10,
		Offset: 0,
	}
	condition := map[string]interface{}{}
	scanResult := entities.Result{
		RepositoryName: "Test 1",
		Status:         "Queued",
		QueuedAt:       currentTime,
		ScanningAt:     currentTime,
		FinishedAt:     currentTime,
	}
	scanResultList := []entities.Result{
		scanResult,
	}
	scanResultResp, _ := adapters.ScanResultRespAdapter(scanResult)
	mockScanResultListResp := []dtos.ScanResultResp{
		scanResultResp,
	}
	expectedScanResultList := dtos.GetScanResultListResp{
		Total:       10,
		ScanResults: mockScanResultListResp,
	}
	mockScanResultRepository := new(MockScanResultRepository)
	mockScanResultRepository.On("FindMany", context.Background(), condition, getScanResultListDto.Limit, getScanResultListDto.Offset).Return(scanResultList, nil)
	mockScanResultRepository.On("Count", context.Background(), condition).Return(10)
	scanResultService := InitScanResultService(mockScanResultRepository)
	scanResultListResp, _ := scanResultService.GetScanResultList(context.Background(), getScanResultListDto)
	assert.Equal(t, expectedScanResultList, scanResultListResp)
	mockScanResultRepository.AssertExpectations(t)
}

func TestDeleteScanResult(t *testing.T) {
	deleteScanResultDto := dtos.DeleteScanResultDto{
		ID: 1,
	}
	condition := map[string]interface{}{
		"id": deleteScanResultDto.ID,
	}
	expectDeleteScanResultResp := dtos.DeleteScanResultResp{
		Message: "Delete Scan Result Success",
	}
	mockScanResultRepository := new(MockScanResultRepository)
	mockScanResultRepository.On("Delete", context.Background(), condition).Return(nil)
	scanResultService := InitScanResultService(mockScanResultRepository)
	deleteScanResultResp, _ := scanResultService.DeleteScanResult(context.Background(), deleteScanResultDto)
	assert.Equal(t, expectDeleteScanResultResp, deleteScanResultResp)
	mockScanResultRepository.AssertExpectations(t)
}

func TestUpdateScanResult(t *testing.T) {
	updateScanResultDto := dtos.UpdateScanResultDto{
		ID:             1,
		RepositoryName: "Test",
		Status:         "Queued",
	}
	condition := map[string]interface{}{
		"id": updateScanResultDto.ID,
	}
	newPayload := entities.Result{
		RepositoryName: updateScanResultDto.RepositoryName,
		Status:         updateScanResultDto.Status,
	}
	expectUpdateScanResultResp := dtos.UpdateScanResultResp{
		Message: "Update Scan Result Success",
	}
	mockScanResultRepository := new(MockScanResultRepository)
	mockScanResultRepository.On("Update", context.Background(), condition, newPayload).Return(nil)
	scanResultService := InitScanResultService(mockScanResultRepository)
	updateScanResultResp, _ := scanResultService.UpdateScanResult(context.Background(), updateScanResultDto)
	assert.Equal(t, expectUpdateScanResultResp, updateScanResultResp)
	mockScanResultRepository.AssertExpectations(t)
}
