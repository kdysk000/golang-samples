package gorm

import (
	"fmt"
	"log"
)

/*
	単一クエリでHasOne(1対1)の関連を持つテーブルのレコードを取得する(Joins)
*/
func GormSample024() {
	fmt.Println("gorm_sample_024")

	db := DbInit()
	AutoMigrate(tables002...)

	records := []Worker{}

	ret := db.Joins("Job").Find(&records)
	if ret.Error != nil {
		log.Fatal(ret.Error)
	}
	for _, worker := range records {
		fmt.Println("ID;", worker.ID, "Name:", worker.Name, "Age:", worker.Age, "Job:", worker.Job.Name)
	}
	fmt.Println("RowsAffected:", ret.RowsAffected)
}
