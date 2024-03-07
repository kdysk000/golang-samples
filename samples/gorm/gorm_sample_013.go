package gorm

import (
	"fmt"
	"log"
)

/*
	条件にマッチする全レコードを取得する(OR句)
*/
func GormSample013() {
	fmt.Println("gorm_sample_013")

	db := DbInit()
	AutoMigrate(tables001...)

	records := []Test001{}

	//Where()の中でOR句を使う
	ret := db.Where("id = 1 OR id = 3").Find(&records)
	if ret.Error != nil {
		log.Fatal(ret.Error)
	}
	for _, record := range records {
		fmt.Println("ID;", record.ID, "UserId:", record.UserId, "Name:", record.Name)
	}
	fmt.Println("RowsAffected:", ret.RowsAffected)

	//Or()メソッドを使う場合
	ret = db.Where("id = 2").Or("id = 4").Find(&records)
	if ret.Error != nil {
		log.Fatal(ret.Error)
	}
	for _, record := range records {
		fmt.Println("ID;", record.ID, "UserId:", record.UserId, "Name:", record.Name)
	}
	fmt.Println("RowsAffected:", ret.RowsAffected)
}

