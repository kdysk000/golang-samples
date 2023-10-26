package myfunc

import (
	"fmt"
)

/*
	関数タイプ変数

	関数も変数で持つことができる
*/
func FuncSample005() {
	fmt.Println("func_sample_005")

	var f func(int) int = testFunc1
	num := f(5)
	fmt.Println(num)
}

func testFunc1(x int) int {
	fmt.Println("This is testFunc1")
	return x
}

/*
  実行結果
  -------
  This is testFunc1
  5
  -------
*/
