package basic

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
func BasicSample007() {
	fmt.Println("basic_sample_007")

	m := map[string]int{
		"apple" : 1,
		"orange": 2,
	}

	fmt.Println(m)

	m["apple"]  = 0  // 存在するキーなので更新
	m["banana"] = 3  // 存在しないキーなので追加

	fmt.Println(m)
}

/*
  実行結果
  -------
  map[apple:1 orange:2]
  map[apple:0 banana:3 orange:2]
  -------
*/
