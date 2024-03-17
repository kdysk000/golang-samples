package gorm

import (
	"fmt"
	"log"
)

/*
	Mapを使ってレコードを作成する

	Mapを使用してレコードを作成する場合、Hookは呼び出されない。
	また、関連テーブルのレコードも保存されず主キーの値も返されない。
*/
func GormSample018() {
	fmt.Println("gorm_sample_018")

	db := DbInit()
	AutoMigrate(tables001...)
	
	ret := db.Model(&User{}).Create([]map[string]interface{}{
		{"Name": RandomStr(7), "Age": RandomInt(2)},
		{"Name": RandomStr(7), "Age": RandomInt(2)},
	})
	if ret.Error != nil {
		log.Fatal(ret.Error)
	}
}

