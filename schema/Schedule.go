package schema

import (
	"time"
)

type Dayofweek string

type Schedule struct {
	Blockend   time.Time `gorm:"column:blockend"`
	Blockstart time.Time `gorm:"column:blockstart"`
	Dayofweek  `gorm:"column:dayofweek"`
	ID         int    `gorm:"column:id;primary_key"`
	MediaType  int    `gorm:"column:media_type"`
	Name       string `gorm:"column:name"`
}

// TableName sets the insert table name for this struct type
func (s *Schedule) TableName() string {
	return "schedule"
}
