package gorm

import (
	"fmt"
	"log"
)

/*
	条件にマッチする全レコードを取得する(Where.Find)
*/
func GormSample010() {
	fmt.Println("gorm_sample_010")

	db := DbInit()
	AutoMigrate(tables001...)

	records := []Test001{}
	ret := db.Where("id < ?", 10).Find(&records)
	if ret.Error != nil {
			log.Fatal(ret.Error)
	}
	for i, record := range records {
		fmt.Println(i, "UserId:", record.UserId, "Name:", record.Name)
	}
	fmt.Println("RowsAffected:", ret.RowsAffected)
}

