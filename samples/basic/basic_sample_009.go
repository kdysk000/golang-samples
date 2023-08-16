package basic

import (
	"fmt"
)

/*
	mapのキーの存在チェック

	書式
	  値, キーの存在有無(bool) := マップ変数[キー]
*/
func BasicSample009() {
	fmt.Println("basic_sample_009")

	m := map[string]int{
		"apple" : 1,
		"orange": 2,
	}

	if num1, ok1 := m["apple"]; ok1 {
		fmt.Println("apple is exist. num:", num1)
	}else{
		fmt.Println("apple is not exist.")
	}
	
	if num2, ok2 := m["banana"]; ok2 {
		fmt.Println("banana is exist. num:", num2)
	}else{
		fmt.Println("banana is not exist.")
	}
}

/*
  実行結果
  -------
  apple is exist. num: 1
  banana is not exist.
  -------
*/
