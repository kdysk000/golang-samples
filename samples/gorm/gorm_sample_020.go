package gorm

import (
	"fmt"
	"log"
)

/*
	取得したレコードの並び替え(Order)
*/
func GormSample020() {
	fmt.Println("gorm_sample_020")

	db := DbInit()
	AutoMigrate(tables001...)

	records := []Test001{}
	//IDの降順に並び替え
	ret := db.Order("id desc").Find(&records)
	if ret.Error != nil {
		log.Fatal(ret.Error)
	}
	for _, record := range records {
		fmt.Println("ID;", record.ID, "UserId:", record.UserId, "Name:", record.Name)
	}
	fmt.Println("RowsAffected:", ret.RowsAffected)
}

