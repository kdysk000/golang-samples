package gorm

import (
	"fmt"
	"log"

	"github.com/google/uuid"
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

	id1 := uuid.New().String()
	id2 := uuid.New().String()
	
	ret := db.Model(&User{}).Create([]map[string]interface{}{
		{"UserId": id1, "Name": "hoge-"+id1},
		{"UserId": id2, "Name": "hoge-"+id2},
	})
	if ret.Error != nil {
		log.Fatal(ret.Error)
	}
}

