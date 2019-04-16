package model

import "github.com/jinzhu/gorm"

type Event struct {
	gorm.Model
	Name string `gorm:"index"`
	Src  string
}
