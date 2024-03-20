package gorm

import (
	"fmt"
	"log"
)

/*
	テーブルにレコードを1つ追加する(Create)
*/
func GormSample001() {
	fmt.Println("gorm_sample_001")

	db := DbInit()
	AutoMigrate(tables001...)
	
	ret := db.Create(&User{
		Name: RandomStr(7),
		Age:  RandomInt(2),
	})
	if ret.Error != nil {
		log.Fatal(ret.Error)
	}
	//ret.RowsAffectedの中には、今回の結果のRow数(行数)が入る
	fmt.Println("count:", ret.RowsAffected)
}

