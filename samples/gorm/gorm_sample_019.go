package gorm

import (
	"fmt"
	"log"
)

/*
	条件にマッチする全レコードを取得する(Not)
*/
func GormSample019() {
	fmt.Println("gorm_sample_019")

	db := DbInit()
	AutoMigrate(tables001...)

	records := []Test001{}

	ret := db.Not("id = 1").Find(&records)
	if ret.Error != nil {
		log.Fatal(ret.Error)
	}
	for _, record := range records {
		fmt.Println("ID;", record.ID, "UserId:", record.UserId, "Name:", record.Name)
	}
	fmt.Println("RowsAffected:", ret.RowsAffected)
}

