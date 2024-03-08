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

	records := []User{}
	ret := db.Where("id < ?", 10).Find(&records)
	if ret.Error != nil {
		log.Fatal(ret.Error)
	}
	for i, record := range records {
		fmt.Println(i, "Name:", record.Name, "Age:", record.Age)
	}
	fmt.Println("RowsAffected:", ret.RowsAffected)
}

