package gorm

import (
	"fmt"
	"log"
)

/*
	単一クエリで関連データを eager loading する(Joins)
*/
func GormSample024() {
	fmt.Println("gorm_sample_024")

	db := DbInit()
	AutoMigrate(tables002...)

	record := []Worker{}

	ret := db.Joins("Job", db.Where(&Job{Name: "Doctor"})).Find(&record)
	if ret.Error != nil {
		log.Fatal(ret.Error)
	}
	fmt.Println(record)
}
