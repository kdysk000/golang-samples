package gorm

import (
	"fmt"
	"log"
)

/*
	条件にマッチする全レコードを取得する(IN句)
*/
func GormSample011() {
	fmt.Println("gorm_sample_011")

	db := DbInit()
	AutoMigrate(tables001...)

	records := []User{}
	ret := db.Where("id IN ?", []string{"2","4","6","8","10"}).Find(&records)
	if ret.Error != nil {
		log.Fatal(ret.Error)
	}
	for _, record := range records {
		fmt.Println("ID;", record.ID, "Name:", record.Name, "Age:", record.Age)
	}
	fmt.Println("RowsAffected:", ret.RowsAffected)
}

