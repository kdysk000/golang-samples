package myfunc

import (
	"fmt"
)

/*
	戻り値なしの関数の定義

	戻り値の定義を省略すれば戻り値なしとなる

	書式：
	  func 関数名 (引数 型) {
	  	// do something
	  	return 値
	  }
*/
func FuncSample003() {
	fmt.Println("func_sample_003")

	printNum(10)
}

func printNum(x int) {
	fmt.Println(x)
}

/*
  実行結果
  -------
  10
  -------
*/
