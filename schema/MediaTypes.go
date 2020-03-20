package schema

import "gopkg.in/guregu/null.v3"

type MediaTypes struct {
	ID          int         `gorm:"column:id;primary_key"`
	Name        null.String `gorm:"column:name"`
	PlayoutType null.String `gorm:"column:playoutType"`
}

// TableName sets the insert table name for this struct type
func (m *MediaTypes) TableName() string {
	return "media_types"
}
