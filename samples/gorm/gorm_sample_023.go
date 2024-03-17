package gorm

import (
	"fmt"
	"log"
)

/*
	取得した結果から重複レコードを1つにまとめる(Distinct)
*/
func GormSample023() {
	fmt.Println("gorm_sample_023")

	db := DbInit()
	AutoMigrate(tables002...)

	res := []Job{}
	ret := db.Distinct("name").Find(&res)
	if ret.Error != nil {
		log.Fatal(ret.Error)
	}
	fmt.Println(res)
}

