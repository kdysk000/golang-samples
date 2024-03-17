package gorm

import (
	"fmt"
	"log"
)

/*
	単一クエリでHasOne(1対1)の関連を持つテーブルのレコードを取得する(Joins)
*/
func GormSample024() {
	fmt.Println("gorm_sample_024")

	db := DbInit()
	AutoMigrate(tables002...)

	record := Worker{}

	ret := db.Joins("Job").First(&record)
	if ret.Error != nil {
		log.Fatal(ret.Error)
	}
	fmt.Println(record)
}
