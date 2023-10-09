package basic

import (
	"fmt"
)

/*
	rangeを使ったfor文（簡易式for文）

	配列の値を順番に取得する場合などに使用

	書式：
		for index, value := range 配列など {
			// 繰り返し処理する内容
		}
*/
func BasicSample014() {
	fmt.Println("basic_sample_014")

	array := [...]string{"foo", "bar", "baz"}

	for i, v := range array {
		fmt.Println(i, v)
	}
}

/*
  実行結果
  -------
  0 foo
  1 bar
  2 baz
  -------
*/
