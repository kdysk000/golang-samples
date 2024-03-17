package gorm

import (
	"fmt"
	"log"
)

/*
	HasOne(1対1)の関連を持つテーブルの作成
*/
func GormSample016() {
	fmt.Println("gorm_sample_016")

	db := DbInit()
	AutoMigrate(tables002...)

	ret := db.Create(&Worker{
		Name: RandomStr(7),
		Age:  RandomInt(2),
		Job:  Job{
			Name: "Teacher",
		},
	})
	if ret.Error != nil {
		log.Fatal(ret.Error)
	}
	//ret.RowsAffectedの中には、今回の結果のRow数(行数)が入る
	fmt.Println("count:", ret.RowsAffected)
}

