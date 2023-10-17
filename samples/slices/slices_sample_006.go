package slices

import (
	"fmt"

	"golang.org/x/exp/slices"
)

/*
	スライスの連続した同じ要素を削除し新しいスライスを返す(CompactFunc)

	func CompactFunc(s S, eq func(E, E) bool) S
	  param:
	    s : スライス
	    eq: 比較関数
	  return:
	    S : 削除後の新しいスライス

	Compactの比較部分を指定できる
*/
func SlicesSample006() {
	fmt.Println("slices_sample_006")

	s1 := []struct {
		Id   string
		Name string
	}{
		{"0001", "hoge"},
		{"0002", "hoge"},
		{"0003", "fuga"},
		{"0003", "fuga"},
	}
	s2 := slices.CompactFunc(s1, func(s, target struct {
		Id   string
		Name string
	}) bool {
		return s.Id == target.Id && s.Name == target.Name
	})
	
	fmt.Println(s1)
	fmt.Println(s2)
}

/*
  実行結果
  -------
  [{0001 hoge} {0002 hoge} {0003 fuga} {0003 fuga}]
  [{0001 hoge} {0002 hoge} {0003 fuga}]
  -------
*/
