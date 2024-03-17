package gorm

import (
	"fmt"
	"log"
)

/*
	HasOne(1対1)の関連先のテーブルを検索条件にしてレコードを取得する(InnerJoins)
*/
func GormSample029() {
	fmt.Println("gorm_sample_029")

	db := DbInit()
	AutoMigrate(tables002...)

	records := []Worker{}

	//Job.Name が Doctor の Workerレコードを全て取得
	ret := db.InnerJoins("Job", db.Where(&Job{Name: "Doctor"})).Find(&records)
	if ret.Error != nil {
		log.Fatal(ret.Error)
	}
	for _, worker := range records {
		fmt.Println("ID;", worker.ID, "Name:", worker.Name, "Age:", worker.Age, "Job:", worker.Job.Name)
	}
	fmt.Println("RowsAffected:", ret.RowsAffected)
}
