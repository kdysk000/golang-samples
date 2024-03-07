package gorm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func DbInit() *gorm.DB {
        var err error
        dsn := "root:secret@tcp(127.0.0.1:3306)/sample_db?charset=utf8mb4&parseTime=true"
        db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
        if err != nil {
                panic("failed to connect database")
        }
		return db
}

func AutoMigrate(models ...interface{}) {
	db.AutoMigrate(models...,)
}

