package myfunc

import (
	"fmt"
)

/*
	関数リテラル(匿名関数)
*/
func FuncSample006() {
	fmt.Println("func_sample_006")

	f := func(n1 int) int {
		fmt.Println("This is testFunc")
		return n1
	}
	num := f(5)
	fmt.Println(num)
}

/*
  実行結果
  -------
  This is testFunc
  5
  -------
*/
