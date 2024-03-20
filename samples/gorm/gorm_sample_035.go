package gorm

import (
	"fmt"
	"log"
)

/*
	ManyToMany(多対多)の関連を持つテーブルの作成
*/
func GormSample035() {
	fmt.Println("gorm_sample_035")

	db := DbInit()
	AutoMigrate(tables005...)

	ret := db.Create(&Speaker{Name: RandomStr(7), Languages: []Language{{Name: "Japanese"}, {Name: "English"}}})
	if ret.Error != nil {
		log.Fatal(ret.Error)
	}
}
