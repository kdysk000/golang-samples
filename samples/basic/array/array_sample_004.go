package array

import (
	"fmt"
)

/*
	配列の要素数を取得(len)
*/
func ArraySample004() {
	fmt.Println("array_sample_004")

	ar := [5]int{1,2,3,4,5}
	num := len(ar)

	fmt.Println(ar)
	fmt.Println(num)
}

/*
  実行結果
  -------
  [1 2 3 4 5]
  5
  -------
*/
