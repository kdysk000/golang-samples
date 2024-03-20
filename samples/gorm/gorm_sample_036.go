package gorm

import (
	"fmt"
	"log"
)

/*
	ManyToMany(多対多)の関連を持つテーブルのレコードを取得する
*/
func GormSample036() {
	fmt.Println("gorm_sample_036")

	db := DbInit()
	AutoMigrate(tables005...)

	ret := db.Create(&Speaker{Name: RandomStr(7), Languages: []Language{{Name: "Japanese"}, {Name: "English"}}})
	if ret.Error != nil {
		log.Fatal(ret.Error)
	}
	speaker := Speaker{}
	ret = db.Preload("Languages").First(&speaker)
	if ret.Error != nil {
		log.Fatal(ret.Error)
	}

	lang := Language{}
	ret = db.Preload("Speakers").First(&lang)
	if ret.Error != nil {
		log.Fatal(ret.Error)
	}

	fmt.Println(speaker)
	fmt.Println(lang)
}
