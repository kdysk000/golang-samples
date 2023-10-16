package slices

import (
	"fmt"

	"golang.org/x/exp/slices"
)

/*
	スライスをソートする(Sort)

	func Sort(x []E)
		param:
		  x: スライス
*/
func SlicesSample016() {
	fmt.Println("slices_sample_016")

	s1 := []int{3, 2, 1}
	slices.Sort(s1)
	fmt.Println(s1)

	s2 := []string{"z", "y", "x"}
	slices.Sort(s2)
	fmt.Println(s2)
}

/*
  実行結果
  -------
  [1 2 3]
  [x y z]
  -------
*/
