package mymap

import (
	"fmt"
)

/*
	mapの宣言と初期化

	mapとは
	　・データ構造の一つで、キーと値を保持する
	　・他の言語で「連想配列」「ディクショナリ」などと呼ばれるもの

	書式
	  変数名 := map[キーのデータ型]値のデータ型{}
*/
func MapSample001() {
	fmt.Println("map_sample_001")

	// mapの定義と初期化
	m := map[string]int{
		"foo": 1,
		"bar": 2,
		"baz": 3,
	}

	fmt.Println(m)
}

/*
  実行結果
  -------
  map[bar:2 baz:3 foo:1]
  -------
*/
