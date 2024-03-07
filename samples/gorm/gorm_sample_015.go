package gorm

import (
	"fmt"
	"log"
)

/*
	構造体で条件を指定してレコードを取得する
*/
func GormSample015() {
	fmt.Println("gorm_sample_015")

	db := DbInit()
	AutoMigrate(tables001...)

	record := Test001{}
	wr := Test001{Name: "hoge"}
	ret := db.Where(&wr).Find(&record)
	if ret.Error != nil {
			log.Fatal(ret.Error)
	}
	fmt.Println(record)
	fmt.Println("RowsAffected:", ret.RowsAffected)
}

