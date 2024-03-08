package gorm

import (
	"fmt"
	"log"
)

/*
	条件にマッチする全レコードを取得する(AND句)
*/
func GormSample012() {
	fmt.Println("gorm_sample_012")

	db := DbInit()
	AutoMigrate(tables001...)

	records := []User{}
	ret := db.Where("id > 0 AND id < 5").Find(&records)
	if ret.Error != nil {
		log.Fatal(ret.Error)
	}
	for _, record := range records {
		fmt.Println("ID;", record.ID, "Name:", record.Name, "Age:", record.Age)
	}
	fmt.Println("RowsAffected:", ret.RowsAffected)
}

