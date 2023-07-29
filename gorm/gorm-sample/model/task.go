package model

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	TaskId    string `json:"taskid"`
	Name      string `json:"name"`
	OwnerID   uint
	OwnerType string
}
