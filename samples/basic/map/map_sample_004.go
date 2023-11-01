package mymap

import (
	"fmt"
)

/*
	mapに要素を追加

	書式
	  変数名[キー] = 値

	・存在するキーの場合は値が更新される
	・存在しないキーの場合は追加となる
*/
func MapSample004() {
	fmt.Println("map_sample_004")

	m := map[string]int{
		"foo": 1,
		"bar": 2,
	}

	fmt.Println(m)

	m["foo"] = 0  // 存在するキーなので更新
	m["baz"] = 3  // 存在しないキーなので追加

	fmt.Println(m)
}

/*
  実行結果
  -------
  map[bar:2 foo:1]
  map[bar:2 baz:3 foo:0]
  -------
*/
