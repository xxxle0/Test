package entities

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Finding struct {
	gorm.Model
	Type     string
	RuleID   string
	ResultId uint
	Location datatypes.JSON
	Metadata datatypes.JSON
}
