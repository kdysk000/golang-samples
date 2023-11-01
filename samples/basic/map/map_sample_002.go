package mymap

import (
	"fmt"
)

/*
	mapから値を取得

	書式
	  変数名[キー]

	  指定のキーの値を取得できる
*/
func MapSample002() {
	fmt.Println("map_sample_002")

	m := map[string]int{
		"foo": 1,
		"bar": 2,
		"baz": 3,
	}

	fmt.Println(m["bar"])
	fmt.Println(m["abc"])  //存在しないキーを指定してもエラーにはならずその型の初期値が返る
}

/*
  実行結果
  -------
  2
  0
  -------
*/
