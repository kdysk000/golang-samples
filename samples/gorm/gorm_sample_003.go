package gorm

import (
	"fmt"
	"log"

	"github.com/google/uuid"
)

/*
	フィールドを選択してレコードを作成する(Select.Create)
*/
func GormSample003() {
	fmt.Println("gorm_sample_003")

	db := DbInit()
	AutoMigrate(tables001...)

	id := uuid.New().String()
	
	ret := db.Select("Name").Create(&Test001{
        UserId: id,  //このフィールドは登録されない
        Name: "hoge-"+id,
	})
	if ret.Error != nil {
			log.Fatal(ret.Error)
	}
}

