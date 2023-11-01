package mymap

import (
	"fmt"
)

/*
	mapの要素を削除する(delete)

	書式
	  delete(map変数, key)
*/
func MapSample005() {
	fmt.Println("map_sample_005")

	m := map[string]int{
		"foo": 1,
		"bar": 2,
		"baz": 3,
	}

	fmt.Println(m)

	delete(m, "baz")

	fmt.Println(m)
}

/*
  実行結果
  -------
  map[bar:2 baz:3 foo:1]
  map[bar:2 foo:1]
  -------
*/
