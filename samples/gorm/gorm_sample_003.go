package gorm

import (
	"fmt"
	"log"
)

/*
	フィールドを選択してレコードを作成する(Select.Create)
*/
func GormSample003() {
	fmt.Println("gorm_sample_003")

	db := DbInit()
	AutoMigrate(tables001...)

	ret := db.Select("Name").Create(&User{
		Name: RandomStr(7),
		Age:  RandomInt(2),  //このフィールドは登録されない
	})
	if ret.Error != nil {
		log.Fatal(ret.Error)
	}
}

