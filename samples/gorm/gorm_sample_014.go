package gorm

import (
	"fmt"
	"log"
)

/*
	条件にマッチする全レコードを取得する(BETWEEN句)
*/
func GormSample014() {
	fmt.Println("gorm_sample_014")

	db := DbInit()
	AutoMigrate(tables001...)

	records := []User{}
	ret := db.Where("id BETWEEN 0 AND 5").Find(&records)
	if ret.Error != nil {
		log.Fatal(ret.Error)
	}
	for _, record := range records {
		fmt.Println("ID;", record.ID, "Name:", record.Name, "Age:", record.Age)
	}
	fmt.Println("RowsAffected:", ret.RowsAffected)
}

