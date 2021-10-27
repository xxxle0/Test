package entities

type Finding struct {
	Type     int
	RuleID   int
	Location FindingLocation
	Metadata Metadata
}

type FindingLocation struct {
	Path      string
	Positions FindingPosition
}

type Metadata struct {
	Description string
	Severity    string
}

type FindingPosition struct {
	Begin struct {
		Line string
	}
}
