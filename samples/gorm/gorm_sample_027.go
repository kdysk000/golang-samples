package gorm

import (
	"fmt"
	"log"
)

/*
	レコードの複数カラムを更新する(Updates)

	注：ゼロ値(数値の0や空文字列)に更新はできない
*/
func GormSample027() {
	fmt.Println("gorm_sample_027")

	db := DbInit()
	AutoMigrate(tables002...)

	worker := Worker{}

	//更新前のレコードを確認
	ret := db.Where("id = ?", 1).First(&worker)
	if ret.Error != nil {
		log.Fatal(ret.Error)
	}
	fmt.Println("before name:", worker.Name, "before age:", worker.Age)

	//更新
	ret = db.Where("id = ?", 1).Updates(Worker{Name: RandomStr(7), Age: RandomInt(2)})
	if ret.Error != nil {
		log.Fatal(ret.Error)
	}

	//更新後のレコードを確認
	ret = db.Where("id = ?", 1).First(&worker)
	if ret.Error != nil {
		log.Fatal(ret.Error)
	}
	fmt.Println("before name:", worker.Name, "before age:", worker.Age)
}
