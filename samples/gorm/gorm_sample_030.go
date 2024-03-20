package gorm

import (
	"fmt"
	"log"
)

/*
	HasMany(1対多)の関連を持つテーブルの作成
*/
func GormSample030() {
	fmt.Println("gorm_sample_030")

	db := DbInit()
	AutoMigrate(tables003...)

	ret := db.Create(&Rich{
		Name: RandomStr(7),
		Age:  RandomInt(2),
		CreditCards:  []CreditCard{
			{Bank: "Mitsui", Number: RandomInt(10)},
			{Bank: "Ufj", Number: RandomInt(10)},
		},
	})
	if ret.Error != nil {
		log.Fatal(ret.Error)
	}

	ret = db.Create(&Rich{
		Name: RandomStr(7),
		Age:  RandomInt(2),
		CreditCards:  []CreditCard{
			{Bank: "Ufj", Number: RandomInt(10)},
			{Bank: "Mizuho", Number: RandomInt(10)},
		},
	})
	if ret.Error != nil {
		log.Fatal(ret.Error)
	}
}

