package gorm

import (
	"fmt"
	"log"
)

/*
	レコードを1つ取得する(Last)

	Last
	  ・プライマリーキーの降順で取得
	  ・プライマリーキーが無い場合は、モデルの最初のフィールドで順序付けされる
*/
func GormSample006() {
	fmt.Println("gorm_sample_006")

	db := DbInit()
	AutoMigrate(tables001...)

	record := Test001{}
	ret := db.Last(&record)
	if ret.Error != nil {
			log.Fatal(ret.Error)
	}
	fmt.Println(record)
}

