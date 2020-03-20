package schema

import "gopkg.in/guregu/null.v3"

type Series struct {
	ID   int         `gorm:"column:id;primary_key"`
	Name null.String `gorm:"column:name"`
	Path null.String `gorm:"column:path"`
}

// TableName sets the insert table name for this struct type
func (s *Series) TableName() string {
	return "series"
}

