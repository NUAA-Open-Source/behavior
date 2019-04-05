package event

import "github.com/jinzhu/gorm"

type Event struct {
	gorm.Model
	Name string `gorm:"index"`
	Src  string
}

type Request struct {
	Name string `json:"name"`
	Src  string `json:"src"`
}
