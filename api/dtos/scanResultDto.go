package dtos

import (
	"time"
)

type CreateScanResultDto struct {
	Status         string       `json:"status"`
	RepositoryName string       `json:"repository_name"`
	Findings       []FindingDto `json:"findings"`
	QueuedAt       time.Time    `json:"queued_at"`
	ScanningAt     time.Time    `json:"scanning_at"`
	FinishedAt     time.Time    `json:"finished_at"`
}

type GetScanResultDetailDto struct {
	ID int `uri:"id"`
}

type GetScanResultListDto struct {
	Limit  int `form:"limit"`
	Offset int `form:"offset"`
}

type DeleteScanResultDto struct {
	ID int `uri:"id"`
}

type UpdateScanResultDto struct {
	ID             int          `uri:"id"`
	Status         string       `json:"status"`
	RepositoryName string       `json:"repository_name"`
	Findings       []FindingDto `json:"findings"`
	QueuedAt       time.Time    `json:"queued_at"`
	ScanningAt     time.Time    `json:"scanning_at"`
	FinishedAt     time.Time    `json:"finished_at"`
}

type CreateScanResultResp struct {
	Message string `json:"message"`
}

type ScanResultResp struct {
	ID             uint         `json:"id"`
	Status         string       `json:"status"`
	RepositoryName string       `json:"repository_name"`
	Findings       []FindingDto `json:"findings,omitempty"`
	QueuedAt       time.Time    `json:"queued_at"`
	ScanningAt     time.Time    `json:"scanning_at"`
	FinishedAt     time.Time    `json:"finished_at"`
}

type GetScanResultDetailResp struct {
	ScanResult ScanResultResp `json:"scan_result"`
}

type GetScanResultListResp struct {
	ScanResults []ScanResultResp `json:"scan_results,omitempty"`
	Total       int64            `json:"total,omitempty"`
}

type DeleteScanResultResp struct {
	Message string `json:"message"`
}

type UpdateScanResultResp struct {
	Message string `json:"message"`
}
