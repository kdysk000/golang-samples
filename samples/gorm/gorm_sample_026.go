package gorm

import (
	"fmt"
	"log"
)

/*
	レコードの単一カラムを更新する(Update)

	Updatesとは違いゼロ値に更新可能
*/
func GormSample026() {
	fmt.Println("gorm_sample_026")

	db := DbInit()
	AutoMigrate(tables002...)

	worker := Worker{}

	//更新前のレコードを確認
	ret := db.Where("id = ?", 1).First(&worker)
	if ret.Error != nil {
		log.Fatal(ret.Error)
	}
	fmt.Println("before name:", worker.Name)

	//更新
	ret = db.Model(&Worker{}).Where("id = ?", 1).Update("age", RandomStr(7))
	if ret.Error != nil {
		log.Fatal(ret.Error)
	}

	//更新後のレコードを確認
	ret = db.Where("id = ?", 1).First(&worker)
	if ret.Error != nil {
		log.Fatal(ret.Error)
	}
	fmt.Println("after name:", worker.Name)
}
