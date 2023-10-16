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

	s1 := []int{1, 2, 3, 4, 5}
	s2 := slices.Clone(s1)
	s1 = append(s1, 6)
	fmt.Println(s1)
	fmt.Println(s2)
}

/*
  実行結果
  -------
  [1 2 3 4 5 6]
  [1 2 3 4 5]
  -------
*/
