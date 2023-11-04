package myfunc

import (
	"fmt"
)

/*
	関数リテラル(無名関数)
*/
func FuncSample008() {
	fmt.Println("func_sample_008")

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
