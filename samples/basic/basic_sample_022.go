package basic

import (
	"fmt"
)

/*
	関数リテラル(匿名関数)

	関数も変数で持つことができる
*/
func BasicSample022() {
	fmt.Println("basic_sample_022")

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
