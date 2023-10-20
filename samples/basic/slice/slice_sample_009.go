package slice

import (
	"fmt"
)

/*
	スライスを空にする

	nilを代入する
*/
func SliceSample009() {
	fmt.Println("slice_sample_009")

	s := []int{1,2,3,4,5}

	fmt.Println(s)

	s = nil
	fmt.Println(s)
}

/*
  実行結果
  -------
  [1 2 3 4 5]
  []
  -------
*/
