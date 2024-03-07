package gorm

import (
	"fmt"
	"log"

	"github.com/google/uuid"
)

/*
	テーブルにレコードを1つ追加する(Create)
*/
func GormSample001() {
	fmt.Println("gorm_sample_001")

	db := DbInit()
	AutoMigrate(tables001...)

	id := uuid.New().String()
	
	ret := db.Create(&Test001{
        UserId: id,
        Name: "hoge",
	})
	if ret.Error != nil {
			log.Fatal(ret.Error)
	}
	//ret.RowsAffectedの中には、今回の結果のRow数(行数)が入る
	fmt.Println("count:", ret.RowsAffected)
}

