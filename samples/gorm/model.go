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
	tables004 = []interface{}{
		&Cat{},
		&Dog{},
		&Toy{},
	}
	tables005 = []interface{}{
		&Speaker{},
		&Language{},
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
	Bank   string
	Number int
	RichID uint
}

// Polymorphism Association
// GORMは Has One と Has Many アソシエーションにおいて、polymorphism associationをサポートしている。
// 所有する側のエンティティのテーブル名が polymorphic type のフィールドに保存され、
// 主キーが polymorphic 用のフィールドに保存される。
type Cat struct {
	ID    int
	Name  string
	Toys  []Toy `gorm:"polymorphic:Owner;"`
}
type Dog struct {
	ID   int
	Name string
	Toys []Toy `gorm:"polymorphic:Owner;"`
}
type Toy struct {
	ID        int
	Name      string
	OwnerID   int     //ここに Cat.ID もしくは Dog.ID が入る
	OwnerType string  //ここに cats もしくは dogs が入る
}

// Many To Many テーブル
// Speaker は複数の言語を所有しかつ言語に属している
// `speaker_languages` は結合テーブルで、GORMの AutoMigrate を使用して
// Speaker テーブルを作成する場合、GORMは自動的に結合テーブルを作成する。
type Speaker struct {
	gorm.Model
	Name       string
	Languages  []Language `gorm:"many2many:speaker_languages;"`
}
type Language struct {
	gorm.Model
	Name       string
	Speakers   []Speaker `gorm:"many2many:speaker_languages;"`
}