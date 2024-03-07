package gorm

import (
	"fmt"
	"log"
)

/*
	条件を指定してレコードを取得する(Where)
*/
func GormSample009() {
	fmt.Println("gorm_sample_009")

	db := DbInit()
	AutoMigrate(tables001...)

	record := Test001{}
	ret := db.Where("id = ?", 2).First(&record)
	if ret.Error != nil {
			log.Fatal(ret.Error)
	}
	fmt.Println(record)
}

