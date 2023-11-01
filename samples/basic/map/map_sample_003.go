package mymap

import (
	"fmt"
)

/*
	make()を用いたmapの作成

	書式
	  変数名 := make(map[key]データ型, キャパシティ)

	  型は必須、キャパシティは任意
*/
func MapSample003() {
	fmt.Println("map_sample_003")

	// mapの定義と初期化
	m := make(map[string]int)

	m["foo"] = 1
	m["bar"] = 2
	m["baz"] = 3

	fmt.Println(m)
}

/*
  実行結果
  -------
  map[bar:2 baz:3 foo:1]
  -------
*/
