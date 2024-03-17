package gorm

import (
	"fmt"
	"log"
)

/*
	グループ化(Group、Having)

	  Group：指定したカラムでグループ化する
	  Having：グループ化した結果に対して絞り込み条件を指定する
*/
type result struct {
	Name  string
	Num   int
}

func GormSample022() {
	fmt.Println("gorm_sample_022")

	db := DbInit()
	AutoMigrate(tables002...)

	//グループ化した結果をそのまま取得
	res := []result{}
	ret := db.Model(&Job{}).Select("name, count(name) as num").Group("name").Find(&res)
	if ret.Error != nil {
		log.Fatal(ret.Error)
	}
	fmt.Println(res)

	//グループ化したあと指定した条件の結果を取得
	ret = db.Model(&Job{}).Select("name, count(name) as num").Group("name").Having("name = ?", "Doctor").Find(&res)
	if ret.Error != nil {
		log.Fatal(ret.Error)
	}
	fmt.Println(res)
}

