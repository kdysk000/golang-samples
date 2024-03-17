package gorm

import (
	"fmt"
	"log"
)

/*
	HasOne(1対1)の関連を持つテーブルのレコードを条件を指定して取得する
*/
func GormSample028() {
	fmt.Println("gorm_sample_028")

	db := DbInit()
	AutoMigrate(tables002...)

	worker := Worker{}

	// WorkerテーブルのIDが 2 のレコードを Jobフィールドも合わせて取得する
	ret := db.Preload("Job").Where("id = ?", 2).First(&worker)
	if ret.Error != nil {
		log.Fatal(ret.Error)
	}
	fmt.Println("ID;", worker.ID, "Name:", worker.Name, "Age:", worker.Age, "Job:", worker.Job.Name, "WorkerID:", worker.Job.WorkerID)
}
