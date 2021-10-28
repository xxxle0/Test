package dtos

type FindingDto struct {
	Type     string                 `json:"type,omitempty"`
	RuleID   string                 `json:"ruleId,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Location map[string]interface{} `json:"location,omitempty"`
}

type FindingLocationDto struct {
	Path      string             `json:"path"`
	Positions FindingPositionDto `json:"positions"`
}

type MetadataDto struct {
	Description string `json:"description"`
	Severity    string `json:"severity"`
}

type FindingPositionDto struct {
	Begin struct {
		Line int `json:"line"`
	} `json:"begin"`
}
