package gorm

import (
	"gorm.io/gorm"
)

var (
	tables001 = []interface{}{
		&Test001{},
	}
	tables002 = []interface{}{
		&Test002{},
		&Test002sub{},
	}
)

type Test001 struct {
        gorm.Model
        UserId  string `json:"userid"`
        Name    string `json:"name"`
}

// Has One テーブル
type Test002 struct {
	gorm.Model
	UserId  string `json:"userid"`
	Name    string `json:"name"`
	Sub     Test002sub `json:"sub"`
}
// Has One を定義する場合、外部キーフィールドも存在する必要がある
// そのフィールドの名前は通常 Has One を持つモデルの型名に primary key を足したものとして作成される
// ここでは Test002 + ID で Test002Id となる
type Test002sub struct {
	gorm.Model
	SubId      string `json:"userid"`
	Test002ID  string `json:"test002id"`  //Test002の主キーがこのフィールドに保存される
}
