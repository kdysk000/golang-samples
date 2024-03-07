package gorm

import (
	"fmt"
	"log"
)

/*
	取得したレコードの並び替え(Order)
*/
func GormSample021() {
	fmt.Println("gorm_sample_021")

	db := DbInit()
	AutoMigrate(tables001...)

	records := []Test001{}
	
	limit := 5  //最大5件
	page := 2   //2ページ目(1ページ5件となるので6~10までのレコードを取得)
	ret := db.Limit(limit).Offset((page-1)*limit).Find(&records)
	if ret.Error != nil {
		log.Fatal(ret.Error)
	}
	for _, record := range records {
		fmt.Println("ID;", record.ID, "UserId:", record.UserId, "Name:", record.Name)
	}
	fmt.Println("RowsAffected:", ret.RowsAffected)
}

