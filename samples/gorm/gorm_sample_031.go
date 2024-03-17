package gorm

import (
	"fmt"
	"log"
)

/*
	HasMany(1対多)の関連を持つテーブルのレコードを取得する(Preload)

	注：HasOneと違いHasManyはPreloadしかEager Load できない？ (Joinsを使うとエラーになる)
*/
func GormSample031() {
	fmt.Println("gorm_sample_031")

	db := DbInit()
	AutoMigrate(tables003...)

	record := Rich{}

	//Preloadなしの場合関連テーブルの情報は取得できないので空
	ret := db.First(&record)
	if ret.Error != nil {
		log.Fatal(ret.Error)
	}
	fmt.Println(record)

	//Preloadで関連テーブルのレコードも取得
	ret = db.Preload("CreditCards").First(&record)
	if ret.Error != nil {
		log.Fatal(ret.Error)
	}
	fmt.Println(record)
}

