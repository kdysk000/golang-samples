package gorm

import (
	"fmt"
	"log"
)

/*
	全てのレコードを取得する(Find)
*/
func GormSample008() {
	fmt.Println("gorm_sample_008")

	db := DbInit()
	AutoMigrate(tables001...)

	records := []Test001{}
	ret := db.Find(&records)
	if ret.Error != nil {
		log.Fatal(ret.Error)
	}
	for i, record := range records {
		fmt.Println(i, "UserId:", record.UserId, "Name:", record.Name)
	}
	fmt.Println("RowsAffected:", ret.RowsAffected)
}

