package gorm

import (
	"fmt"
	"log"
)

/*
	レコードの特定カラムだけを別モデルに格納する(Scan)
*/

type Result struct {
	Name string
	Age  int
}

func GormSample025() {
	fmt.Println("gorm_sample_025")

	db := DbInit()
	AutoMigrate(tables002...)

	var result Result

	ret := db.Table("workers").Select("name", "age").Where("id = ?", 1).Scan(&result)
	if ret.Error != nil {
		log.Fatal(ret.Error)
	}
	fmt.Println(result)
}
