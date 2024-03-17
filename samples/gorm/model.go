package gorm

import (
	"gorm.io/gorm"
)

var (
	tables001 = []interface{}{
		&User{},
	}
	tables002 = []interface{}{
		&Worker{},
		&Job{},
	}
	tables003 = []interface{}{
		&Rich{},
		&CreditCard{},
	}
)

type User struct {
	gorm.Model
	Name   string `json:"name"`
	Age    int `json:"age"`
}

// Has One テーブル
type Worker struct {
	gorm.Model
	Name   string `json:"name"`
	Age    int `json:"age"`
	Job    Job `json:"job"`
}
// Has One を定義する場合、外部キーフィールドも存在する必要がある
// そのフィールドの名前は通常 Has One を持つモデルの型名に primary key を足したものとして作成される
// ここでは Worker + ID で WorkerID となる
type Job struct {
	gorm.Model
	Name      string `json:"jobid"`
	WorkerID  string `json:"workerid"`  //Workerの主キーがこのフィールドに保存される
}

// Has Many テーブル
type Rich struct {
	gorm.Model
	Name        string `json:"name"`
	Age         int `json:"age"`
	CreditCards []CreditCard `json:"credit_card"`
}
// Has Many を定義する場合、外部キーフィールドも存在する必要がある
// そのフィールドの名前は通常 Has Many を持つモデルの型名に primary key を足したものとして作成される
// ここでは Rich + ID で RichID となる
type CreditCard struct {
	gorm.Model
	Number int
	RichID uint
  }
