package gorm

import (
	"fmt"
	"log"
)

/*
	HasOne(1対1)の関連を持つテーブルのレコードを取得する(Preload)
*/
func GormSample017() {
	fmt.Println("gorm_sample_017")

	db := DbInit()
	AutoMigrate(tables002...)

	record := Worker{}

	//Preloadなしの場合関連テーブルの情報は取得できないので空
	ret := db.First(&record)
	if ret.Error != nil {
		log.Fatal(ret.Error)
	}
	fmt.Println(record)

	//Preloadで関連テーブルのレコードも取得
	ret = db.Preload("Job").First(&record)
	if ret.Error != nil {
		log.Fatal(ret.Error)
	}
	fmt.Println(record)
}

