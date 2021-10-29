package adapters

import (
	"testing"
	"time"

	"github.com/duybkit13/api/dtos"
	"github.com/duybkit13/api/entities"
	"github.com/stretchr/testify/assert"
)

func InitTableDrivenTest() {

}

func TestScanResultRespAdapter(t *testing.T) {
	currentTime := time.Now()
	scanResult := entities.Result{
		RepositoryName: "Test Repo 1",
		Status:         "Success",
		QueuedAt:       currentTime,
		FinishedAt:     currentTime,
		ScanningAt:     currentTime,
	}
	exectedScanResultResp := dtos.ScanResultResp{
		RepositoryName: "Test Repo 1",
		Status:         "Success",
		QueuedAt:       currentTime,
		FinishedAt:     currentTime,
		ScanningAt:     currentTime,
	}
	scanResultResp, _ := ScanResultRespAdapter(scanResult)
	assert.Equal(t, scanResultResp.Status, exectedScanResultResp.Status, "got %q, wanted %q")
	assert.Equal(t, scanResultResp.RepositoryName, exectedScanResultResp.RepositoryName, "got %q, wanted %q")
	assert.Equal(t, scanResultResp.QueuedAt, exectedScanResultResp.QueuedAt, "got %q, wanted %q")
	assert.Equal(t, scanResultResp.ScanningAt, exectedScanResultResp.ScanningAt, "got %q, wanted %q")
	assert.Equal(t, scanResultResp.FinishedAt, exectedScanResultResp.FinishedAt, "got %q, wanted %q")
}

func TestScanResultAdapter(t *testing.T) {
	currentTime := time.Now()
	scanResultDto := dtos.CreateScanResultDto{
		RepositoryName: "Test Repo",
		Status:         "Fail",
		QueuedAt:       currentTime,
		FinishedAt:     currentTime,
		ScanningAt:     currentTime,
	}
	scanResult := entities.Result{
		RepositoryName: "Test Repo",
		Status:         "Fail",
		QueuedAt:       currentTime,
		FinishedAt:     currentTime,
		ScanningAt:     currentTime,
	}
	expectedScanResult, _ := ScanResultAdapter(scanResultDto)
	assert.Equal(t, scanResult.Status, expectedScanResult.Status, "got %q, wanted %q")
	assert.Equal(t, scanResult.RepositoryName, expectedScanResult.RepositoryName, "got %q, wanted %q")
	assert.Equal(t, scanResult.QueuedAt, expectedScanResult.QueuedAt, "got %q, wanted %q")
	assert.Equal(t, scanResult.ScanningAt, expectedScanResult.ScanningAt, "got %q, wanted %q")
	assert.Equal(t, scanResult.FinishedAt, expectedScanResult.FinishedAt, "got %q, wanted %q")
}
