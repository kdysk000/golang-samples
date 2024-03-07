package gorm

import (
	"fmt"
	"log"

	"github.com/google/uuid"
)

/*
	HasOne(1対1)の関連を持つテーブルの作成
*/
func GormSample016() {
	fmt.Println("gorm_sample_016")

	db := DbInit()
	AutoMigrate(tables002...)

	ret := db.Create(&Test002{
        UserId: uuid.New().String(),
        Name: "hoge",
		Sub:  Test002sub{
			SubId: uuid.New().String(),
		},
	})
	if ret.Error != nil {
			log.Fatal(ret.Error)
	}
	//ret.RowsAffectedの中には、今回の結果のRow数(行数)が入る
	fmt.Println("count:", ret.RowsAffected)
}

