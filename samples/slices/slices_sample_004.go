package slices

import (
	"fmt"

	"golang.org/x/exp/slices"
)

/*
	スライスのコピー(Clone)

	func Clone(s S) S
		param:
		  s: スライス
		return:
		  S: コピーしたスライス
*/
func SlicesSample004() {
	fmt.Println("slices_sample_004")

	x := []int{1, 2, 3, 4, 5}
	y := slices.Clone(x)
	x = append(x, 6)
	fmt.Println(x)
	fmt.Println(y)
}

/*
  実行結果
  -------
  [1 2 3 4 5 6]
  [1 2 3 4 5]
  -------
*/
