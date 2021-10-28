package adapters

import (
	"encoding/json"

	"github.com/duybkit13/api/dtos"
	"github.com/duybkit13/api/entities"
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
	scanResultResp.Status = scanResult.Status
	scanResultResp.QueuedAt = scanResult.QueuedAt
	scanResultResp.ScanningAt = scanResult.ScanningAt
	scanResultResp.FinishedAt = scanResult.FinishedAt
	scanResultResp.RepositoryName = scanResult.RepositoryName
	scanResultResp.Findings = scanResultFindingResp
	return scanResultResp, nil
}
