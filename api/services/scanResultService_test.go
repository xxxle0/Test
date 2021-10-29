package services

import (
	"context"
	"testing"

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
	return ret.Error(1)
}

func (m *MockScanResultRepository) FindOne(ctx context.Context, condition map[string]interface{}) (entities.Result, error) {
	ret := m.Called(ctx, condition)
	return ret.Get(0).(entities.Result), ret.Error(1)
}

func (m *MockScanResultRepository) FindMany(ctx context.Context, condition map[string]interface{}, limit, offset int) ([]entities.Result, error) {
	ret := m.Called(ctx, condition)
	return ret.Get(0).([]entities.Result), ret.Error(1)
}

func (m *MockScanResultRepository) Delete(ctx context.Context, condition map[string]interface{}) error {
	ret := m.Called(ctx, condition)
	return ret.Error(1)
}

func (m *MockScanResultRepository) Update(ctx context.Context, condition map[string]interface{}, newPayload entities.Result) error {
	ret := m.Called(ctx, condition, newPayload)
	return ret.Error(1)
}

func (m *MockScanResultRepository) Count(ctx context.Context, condition map[string]interface{}) int64 {
	ret := m.Called(ctx, condition)
	return int64(ret.Int(0))
}

func TestCreateScanResult(t *testing.T) {
	mockScanResult := entities.Result{
		Status:         "Queued",
		RepositoryName: "Repo Test 1",
	}
	createScanResultDto := dtos.CreateScanResultDto{
		Status:         "Queued",
		RepositoryName: "Repo Test 1",
	}
	expectCreateScanResultResp := dtos.CreateScanResultResp{}
	mockScanResultRepository := new(MockScanResultRepository)
	mockScanResultRepository.On("Create", context.Background(), mockScanResult).Return(nil)
	scanResultService := InitScanResultService(mockScanResultRepository)
	createScanResultResp, _ := scanResultService.CreateScanResult(context.Background(), createScanResultDto)
	assert.Equal(t, expectCreateScanResultResp, createScanResultResp)
	mockScanResultRepository.AssertExpectations(t)
}

func TestFindOneScanResult(t *testing.T) {
	condition := map[string]interface{}{
		"id": 1,
	}
	mockScanResult := entities.Result{
		Status:         "Queued",
		RepositoryName: "Repo Test 1",
	}
	scanResultDto := dtos.GetScanResultDetailDto{
		ID: 1,
	}
	expectScanResultResp := dtos.GetScanResultDetailResp{
		ScanResult: dtos.ScanResultResp{},
	}
	mockScanResultRepository := new(MockScanResultRepository)
	mockScanResultRepository.On("FindOne", context.Background(), condition).Return(mockScanResult)
	scanResultService := InitScanResultService(mockScanResultRepository)
	scanResultResp, _ := scanResultService.GetScanResultDetail(context.Background(), scanResultDto)
	assert.Equal(t, expectScanResultResp, scanResultResp, "got %q, wanted %q")
	mockScanResultRepository.AssertExpectations(t)
}

func TestFindManyScanResult(t *testing.T) {
	mockScanResultRepository := new(MockScanResultRepository)
	mockScanResultRepository.On("FindMany", context.Background()).Return(nil)
	mockScanResultRepository.AssertExpectations(t)
}

func TestDeleteScanResult(t *testing.T) {
	condition := map[string]interface{}{}
	deleteScanResultDto := dtos.DeleteScanResultDto{}
	expectDeleteScanResultResp := dtos.UpdateScanResultResp{
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
	condition := map[string]interface{}{}
	updateScanResultDto := dtos.UpdateScanResultDto{}
	expectUpdateScanResultResp := dtos.UpdateScanResultResp{
		Message: "Update Scan Result Success",
	}
	mockScanResultRepository := new(MockScanResultRepository)
	mockScanResultRepository.On("Update", context.Background(), condition).Return(nil)
	scanResultService := InitScanResultService(mockScanResultRepository)
	updateScanResultResp, _ := scanResultService.UpdateScanResult(context.Background(), updateScanResultDto)
	assert.Equal(t, expectUpdateScanResultResp, updateScanResultResp)
	mockScanResultRepository.AssertExpectations(t)
}
