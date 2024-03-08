package gorm

import (
	"fmt"
	"log"
)

/*
	省略するフィールドを指定してレコードを作成する(Omit.Create)
*/
func GormSample004() {
	fmt.Println("gorm_sample_004")

	db := DbInit()
	AutoMigrate(tables001...)
	
	ret := db.Omit("Age").Create(&User{
		Name: RandomStr(7),
		Age:  RandomInt(2),  //このフィールドは登録されない
	})
	if ret.Error != nil {
		log.Fatal(ret.Error)
	}
}

