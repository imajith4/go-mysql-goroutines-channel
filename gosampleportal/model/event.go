package model

// Event model
type Event struct {
	ID     int    `gorm:"type:int"`
	Name   string `gorm:"type:varchar(250)"`
	Userid int64  `gorm:"type:int"`
}
