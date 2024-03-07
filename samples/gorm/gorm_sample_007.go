package gorm

import (
	"fmt"
	"log"
)

/*
	レコードを1つ取得する(Take)

	Take
	  ・特に条件を指定せず取得
*/
func GormSample007() {
	fmt.Println("gorm_sample_007")

	db := DbInit()
	AutoMigrate(tables001...)

	record := Test001{}
	ret := db.Take(&record)
	if ret.Error != nil {
		log.Fatal(ret.Error)
	}
	fmt.Println(record)
}

