package mymap

import (
	"fmt"
)

/*
	mapのキーの存在チェック

	書式
	  値, キーの存在有無(bool) := マップ変数[キー]
*/
func MapSample006() {
	fmt.Println("map_sample_006")

	m := map[string]int{
		"foo": 1,
		"bar": 2,
	}

	if num1, ok1 := m["foo"]; ok1 {
		fmt.Println("foo is exist. num:", num1)
	}else{
		fmt.Println("foo is not exist.")
	}
	
	if num2, ok2 := m["baz"]; ok2 {
		fmt.Println("baz is exist. num:", num2)
	}else{
		fmt.Println("baz is not exist.")
	}
}

/*
  実行結果
  -------
  foo is exist. num: 1
  baz is not exist.
  -------
*/
