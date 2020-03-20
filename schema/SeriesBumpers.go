package schema

import "gopkg.in/guregu/null.v3"

type SeriesBumpers struct {
	ID      int         `gorm:"column:id;primary_key"`
	Inpath  null.String `gorm:"column:inpath"`
	Outpath null.String `gorm:"column:outpath"`
	Series  int         `gorm:"column:series"`
}

// TableName sets the insert table name for this struct type
func (s *SeriesBumpers) TableName() string {
	return "series_bumpers"
}

