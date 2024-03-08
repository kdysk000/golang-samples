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

	records := []User{}
	ret := db.Find(&records)
	if ret.Error != nil {
		log.Fatal(ret.Error)
	}
	for i, record := range records {
		fmt.Println(i, "Name:", record.Name, "Age:", record.Age)
	}
	fmt.Println("RowsAffected:", ret.RowsAffected)
}

