package gorm

import (
	"fmt"
	"log"
)

/*
	Polymorphismの関連を持つテーブルのレコードを取得する
*/
func GormSample034() {
	fmt.Println("gorm_sample_034")

	db := DbInit()
	AutoMigrate(tables004...)

	dog := Dog{}
	ret := db.Preload("Toys").First(&dog)
	if ret.Error != nil {
		log.Fatal(ret.Error)
	}

	cat := Cat{}
	ret = db.Preload("Toys").First(&cat)
	if ret.Error != nil {
		log.Fatal(ret.Error)
	}

	fmt.Println(dog)
	fmt.Println(cat)
}

