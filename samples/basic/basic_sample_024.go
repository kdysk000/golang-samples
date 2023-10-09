package basic

import (
	"fmt"
)

/*
	関数を引数にとる関数
*/
func BasicSample024() {
	fmt.Println("basic_sample_024")

	num := testFunc2(testFunc1)
	fmt.Println(num)
}

func testFunc2(f func(int) int) int {
    fmt.Println("This is testFunc2")
	return f(5)
}

/*
  実行結果
  -------
  This is testFunc2
  This is testFunc1
  5
  -------
*/
