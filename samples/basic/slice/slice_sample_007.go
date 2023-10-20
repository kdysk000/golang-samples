package slice

import (
	"fmt"
)

/*
	スライス特定の要素を取り出す
*/
func SliceSample007() {
	fmt.Println("slice_sample_007")

	s := []int{1,2,3,4,5}

	fmt.Println(s[0])          //先頭
	fmt.Println(s[len(s)-1])   //末尾
}

/*
  実行結果
  -------
  1
  5
  -------
*/
