package gorm

import (
	"fmt"
	"log"
)

/*
	ページネーションでのレコード取得(Limit、Offset)

	  Limit ：取得するレコード数の最大値を指定
	  Offset：取得レコードの先頭いくつをスキップするかを指定
*/
func GormSample021() {
	fmt.Println("gorm_sample_021")

	db := DbInit()
	AutoMigrate(tables001...)

	records := []User{}
	
	limit := 5  //最大5件
	page := 2   //2ページ目
	//1ページ5件となるので6~10までのレコードを取得
	ret := db.Limit(limit).Offset((page-1)*limit).Find(&records)
	if ret.Error != nil {
		log.Fatal(ret.Error)
	}
	for _, record := range records {
		fmt.Println("ID;", record.ID, "Name:", record.Name, "Age:", record.Age)
	}
	fmt.Println("RowsAffected:", ret.RowsAffected)
}

