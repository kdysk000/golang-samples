package basic

import (
	"fmt"
)

/*
	関数を戻り値として返す関数
*/
func BasicSample025() {
	fmt.Println("basic_sample_025")

	f := testFunc3()
	fmt.Println(f(5))
}

func testFunc3() func(int) int {
	fmt.Println("This is testFunc3")

	f := func(x int) int {
		return x
	}
	return f
}

/*
  実行結果
  -------
  This is testFunc3
  5
  -------
*/
