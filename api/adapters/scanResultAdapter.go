package adapters

import (
	"encoding/json"

	"github.com/duybkit13/api/dtos"
	"github.com/duybkit13/api/entities"
	"gorm.io/datatypes"
)

func ScanResultRespAdapter(scanResult entities.Result) (dtos.ScanResultResp, error) {
	var scanResultResp dtos.ScanResultResp
	scanResultFindingResp := make([]dtos.FindingDto, len(scanResult.Findings))
	for i, v := range scanResult.Findings {
		var location map[string]interface{}
		var metadata map[string]interface{}
		locationByteSlice, err := json.Marshal(v.Location)
		if err != nil {
			return scanResultResp, err
		}
		err = json.Unmarshal(locationByteSlice, &location)
		if err != nil {
			return scanResultResp, err
		}
		metadataByteSlice, err := json.Marshal(v.Metadata)
		if err != nil {
			return scanResultResp, err
		}
		err = json.Unmarshal(metadataByteSlice, &metadata)
		if err != nil {
			return scanResultResp, err
		}
		scanResultFindingResp[i].Type = v.Type
		scanResultFindingResp[i].RuleID = v.RuleID
		scanResultFindingResp[i].Location = location
		scanResultFindingResp[i].Metadata = metadata
	}
	scanResultResp.ID = scanResult.ID
	scanResultResp.Status = scanResult.Status
	scanResultResp.QueuedAt = scanResult.QueuedAt
	scanResultResp.ScanningAt = scanResult.ScanningAt
	scanResultResp.FinishedAt = scanResult.FinishedAt
	scanResultResp.RepositoryName = scanResult.RepositoryName
	scanResultResp.Findings = scanResultFindingResp
	return scanResultResp, nil
}

func ScanResultAdapter(scanResultDto dtos.CreateScanResultDto) (entities.Result, error) {
	findings := make([]entities.Finding, len(scanResultDto.Findings))
	for i, v := range scanResultDto.Findings {
		locationByteSlice, err := json.Marshal(v.Location)
		if err != nil {
			return entities.Result{}, err
		}
		metadataByteSlice, err := json.Marshal(v.Metadata)
		if err != nil {
			return entities.Result{}, err
		}
		findings[i] = entities.Finding{
			Type:     v.Type,
			RuleID:   v.RuleID,
			Location: datatypes.JSON(locationByteSlice),
			Metadata: datatypes.JSON(metadataByteSlice),
		}
	}
	scanResult := entities.Result{
		Status:         scanResultDto.Status,
		RepositoryName: scanResultDto.RepositoryName,
		ScanningAt:     scanResultDto.ScanningAt,
		QueuedAt:       scanResultDto.QueuedAt,
		FinishedAt:     scanResultDto.FinishedAt,
		Findings:       findings,
	}
	return scanResult, nil
}
