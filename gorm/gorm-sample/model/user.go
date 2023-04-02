package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserId  string `json:"userid"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}
