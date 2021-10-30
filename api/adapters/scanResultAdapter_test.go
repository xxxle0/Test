package adapters

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/duybkit13/api/dtos"
	"github.com/duybkit13/api/entities"
	"github.com/stretchr/testify/assert"
	"gorm.io/datatypes"
)

func TestScanResultRespAdapter(t *testing.T) {
	currentTime := time.Now()
	metadata := map[string]interface{}{
		"description": "TLS InsecureSkipVerify set true.",
		"severity":    "HIGH",
	}
	location := map[string]interface{}{
		"path": "connectors/apigateway.go",
		"positions": map[string]interface{}{
			"begin": map[string]interface{}{
				"line": float64(60),
			},
		},
	}
	metadataByteSlice, _ := json.Marshal(metadata)
	locationByteSlice, _ := json.Marshal(location)
	scanResult := entities.Result{
		RepositoryName: "Test Repo 1",
		Status:         "Success",
		QueuedAt:       currentTime,
		FinishedAt:     currentTime,
		ScanningAt:     currentTime,
		Findings: []entities.Finding{
			entities.Finding{
				Type:     "sast",
				RuleID:   "G402",
				Location: datatypes.JSON(locationByteSlice),
				Metadata: datatypes.JSON(metadataByteSlice),
			},
		},
	}
	exectedScanResultResp := dtos.ScanResultResp{
		RepositoryName: "Test Repo 1",
		Status:         "Success",
		QueuedAt:       currentTime,
		FinishedAt:     currentTime,
		ScanningAt:     currentTime,
		Findings: []dtos.FindingDto{
			dtos.FindingDto{
				Type:     "sast",
				RuleID:   "G402",
				Metadata: metadata,
				Location: location,
			},
		},
	}
	wrongScanResultResp := dtos.ScanResultResp{
		RepositoryName: "Test Repo 1",
		Status:         "Success",
		QueuedAt:       currentTime,
		FinishedAt:     currentTime,
		ScanningAt:     currentTime,
		Findings: []dtos.FindingDto{
			dtos.FindingDto{
				Type:     "1",
				RuleID:   "G402",
				Metadata: metadata,
				Location: location,
			},
		},
	}
	scanResultResp, _ := ScanResultRespAdapter(scanResult)
	assert.Equal(t, scanResultResp, exectedScanResultResp)
	assert.NotEqual(t, scanResultResp, wrongScanResultResp)
}

func TestScanResultAdapter(t *testing.T) {
	currentTime := time.Now()
	metadata := map[string]interface{}{
		"description": "TLS InsecureSkipVerify set true.",
		"severity":    "HIGH",
	}
	location := map[string]interface{}{
		"path": "connectors/apigateway.go",
		"positions": map[string]interface{}{
			"begin": map[string]interface{}{
				"line": 60,
			},
		},
	}
	metadataByteSlice, _ := json.Marshal(metadata)
	locationByteSlice, _ := json.Marshal(location)
	scanResultDto := dtos.CreateScanResultDto{
		RepositoryName: "Test Repo",
		Status:         "Fail",
		QueuedAt:       currentTime,
		FinishedAt:     currentTime,
		ScanningAt:     currentTime,
		Findings: []dtos.FindingDto{
			dtos.FindingDto{
				Type:     "sast",
				RuleID:   "G402",
				Metadata: metadata,
				Location: location,
			},
		},
	}
	expectScanResult := entities.Result{
		RepositoryName: "Test Repo",
		Status:         "Fail",
		QueuedAt:       currentTime,
		FinishedAt:     currentTime,
		ScanningAt:     currentTime,
		Findings: []entities.Finding{
			entities.Finding{
				Type:     "sast",
				RuleID:   "G402",
				Metadata: datatypes.JSON(metadataByteSlice),
				Location: datatypes.JSON(locationByteSlice),
			},
		},
	}
	wrongExpectScanResult := entities.Result{
		RepositoryName: "Test Repo",
		QueuedAt:       currentTime,
		FinishedAt:     currentTime,
		ScanningAt:     currentTime,
		Findings: []entities.Finding{
			entities.Finding{
				Type:   "test",
				RuleID: "test",
			},
		},
	}
	scanResult, _ := ScanResultAdapter(scanResultDto)
	assert.Equal(t, expectScanResult, scanResult)
	assert.NotEqual(t, wrongExpectScanResult, scanResult)
}
