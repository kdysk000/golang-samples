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

	records := []Test001{}
	ret := db.Where("id > 0 AND id < 5").Find(&records)
	if ret.Error != nil {
			log.Fatal(ret.Error)
	}
	for _, record := range records {
		fmt.Println("ID;", record.ID, "UserId:", record.UserId, "Name:", record.Name)
	}
	fmt.Println("RowsAffected:", ret.RowsAffected)
}

