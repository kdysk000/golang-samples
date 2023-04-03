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
	Tasks   []Task `gorm:"polymorphic:Owner;"`
}

type Task struct {
	gorm.Model
	TaskId    string `json:"taskid"`
	Name      string `json:"name"`
	OwnerID   uint
	OwnerType string
}
