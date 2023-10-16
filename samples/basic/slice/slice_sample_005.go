package slice

import (
	"fmt"
)

/*
	スライスの要素数を取得(len)
*/
func SliceSample005() {
	fmt.Println("slice_sample_005")

	s := []int{1,2,3,4,5}
	num := len(s)

	fmt.Println(s)
	fmt.Println(num)
}

/*
  実行結果
  -------
  [1 2 3 4 5]
  5
  -------
*/
