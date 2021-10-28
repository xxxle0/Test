package adapters

import (
	"testing"
	"time"

	"github.com/duybkit13/api/dtos"
	"github.com/duybkit13/api/entities"
)

func TestScanResultRespAdapter(t *testing.T) {
	currentTime := time.Now()
	scanResult := entities.Result{
		Status:         "Success",
		RepositoryName: "Test",
		QueuedAt:       currentTime,
		FinishedAt:     currentTime,
		ScanningAt:     currentTime,
	}
	exectedScanResultResp := dtos.ScanResultResp{
		Status:         "Success",
		RepositoryName: "Test",
		QueuedAt:       currentTime,
		FinishedAt:     currentTime,
		ScanningAt:     currentTime,
	}
	scanResultResp, _ := ScanResultRespAdapter(scanResult)
	if scanResultResp.Status != exectedScanResultResp.Status {
		t.Errorf("got %q, wanted %q", scanResultResp.Status, exectedScanResultResp.Status)
	}
	if scanResultResp.RepositoryName != exectedScanResultResp.RepositoryName {
		t.Errorf("got %q, wanted %q", scanResultResp.RepositoryName, exectedScanResultResp.RepositoryName)
	}
	if scanResultResp.QueuedAt != exectedScanResultResp.QueuedAt {
		t.Errorf("got %q, wanted %q", scanResultResp.QueuedAt, exectedScanResultResp.QueuedAt)
	}
	if scanResultResp.ScanningAt != exectedScanResultResp.ScanningAt {
		t.Errorf("got %q, wanted %q", scanResultResp.ScanningAt, exectedScanResultResp.ScanningAt)
	}
	if scanResultResp.FinishedAt != exectedScanResultResp.FinishedAt {
		t.Errorf("got %q, wanted %q", scanResultResp.FinishedAt, exectedScanResultResp.FinishedAt)
	}
}

func TestScanResultAdapter(t *testing.T) {
	currentTime := time.Now()
	scanResultDto := dtos.CreateScanResultDto{
		RepositoryName: "test",
		Status:         "Fail",
		QueuedAt:       currentTime,
		FinishedAt:     currentTime,
		ScanningAt:     currentTime,
	}
	scanResult := entities.Result{
		RepositoryName: "test",
		Status:         "Fail",
		QueuedAt:       currentTime,
		FinishedAt:     currentTime,
		ScanningAt:     currentTime,
	}
	expectedScanResult, _ := ScanResultAdapter(scanResultDto)
	if scanResult.RepositoryName != expectedScanResult.RepositoryName {
		t.Errorf("got %q, wanted %q", scanResult.RepositoryName, expectedScanResult.RepositoryName)
	}
	if scanResult.Status != expectedScanResult.Status {
		t.Errorf("got %q, wanted %q", scanResult.Status, expectedScanResult.Status)
	}
	if scanResult.ScanningAt != expectedScanResult.ScanningAt {
		t.Errorf("got %q, wanted %q", scanResult.ScanningAt, expectedScanResult.ScanningAt)
	}
	if scanResult.FinishedAt != expectedScanResult.FinishedAt {
		t.Errorf("got %q, wanted %q", scanResult.FinishedAt, expectedScanResult.FinishedAt)
	}
	if scanResult.QueuedAt != expectedScanResult.QueuedAt {
		t.Errorf("got %q, wanted %q", scanResult.QueuedAt, expectedScanResult.QueuedAt)
	}
}
