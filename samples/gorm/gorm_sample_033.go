package gorm

import (
	"fmt"
	"log"
)

/*
	Polymorphismの関連を持つテーブルの作成
*/
func GormSample033() {
	fmt.Println("gorm_sample_033")

	db := DbInit()
	AutoMigrate(tables004...)

	ret := db.Create(&Dog{Name: RandomStr(7), Toys: []Toy{{Name: RandomStr(7)}, {Name: RandomStr(7)}}})
	if ret.Error != nil {
		log.Fatal(ret.Error)
	}
	ret = db.Create(&Cat{Name: RandomStr(7), Toys: []Toy{{Name: RandomStr(7)}, {Name: RandomStr(7)}}})
	if ret.Error != nil {
		log.Fatal(ret.Error)
	}
}

