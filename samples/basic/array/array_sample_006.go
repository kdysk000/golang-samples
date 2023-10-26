package array

import (
	"fmt"
)

/*
	配列の要素数を取得(len)
*/
func ArraySample006() {
	fmt.Println("array_sample_006")

	ar := [5]int{1,2,3,4,5}
	num := len(ar)

	fmt.Println("ar:", ar)
	fmt.Println("len(ar):", num)
}

/*
  実行結果
  -------
  ar: [1 2 3 4 5]
  len(ar): 5
  -------
*/
