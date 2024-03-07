package gorm

import (
	"fmt"
	"log"

	"github.com/google/uuid"
)

/*
	省略するフィールドを指定してレコードを作成する(Omit.Create)
*/
func GormSample004() {
	fmt.Println("gorm_sample_004")

	db := DbInit()
	AutoMigrate(tables001...)

	id := uuid.New().String()
	
	ret := db.Omit("Name").Create(&Test001{
        UserId: id,
        Name: "hoge-"+id,  //このフィールドは登録されない
	})
	if ret.Error != nil {
			log.Fatal(ret.Error)
	}
}

