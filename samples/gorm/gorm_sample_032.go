package gorm

import (
	"fmt"
	"log"
)

/*
	HasMany(1対多)の関連を持つテーブルのレコードを条件を指定して取得する
*/
func GormSample032() {
	fmt.Println("gorm_sample_032")

	db := DbInit()
	AutoMigrate(tables003...)

	record := Rich{}

	ret := db.Preload("CreditCards").Where("id = ?", 1).First(&record)
	if ret.Error != nil {
		log.Fatal(ret.Error)
	}
	fmt.Println(record)
}
