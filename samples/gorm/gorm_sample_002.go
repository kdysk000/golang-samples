package gorm

import (
	"fmt"
	"log"

	"github.com/google/uuid"
)

/*
	テーブルに複数レコードを追加する(Create)

	Create()にモデルのスライスを渡すことで複数レコードを一括登録できる
*/
func GormSample002() {
	fmt.Println("gorm_sample_002")

	db := DbInit()
	AutoMigrate(tables001...)

	records := []Test001{}

	for i := 0; i < 10; i++ {
		id := uuid.New().String()
		record := Test001{
			UserId: id,
			Name: "hoge-"+id,
		}
		records = append(records, record)
	}
	ret := db.Create(&records)
	if ret.Error != nil {
			log.Fatal(ret.Error)
	}
	//ret.RowsAffectedの中には、今回の結果のRow数(行数)が入る
	fmt.Println("count:", ret.RowsAffected)
}

