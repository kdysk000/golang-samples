package gorm

import (
	"fmt"
	"log"
)

/*
	レコードを1つ取得する(First)

	First
	  ・プライマリーキーの昇順で取得
	  ・プライマリーキーが無い場合は、モデルの最初のフィールドで順序付けされる
*/
func GormSample005() {
	fmt.Println("gorm_sample_005")

	db := DbInit()
	AutoMigrate(tables001...)

	record := Test001{}
	ret := db.First(&record)
	if ret.Error != nil {
			log.Fatal(ret.Error)
	}
	fmt.Println(record)
}

