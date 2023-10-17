package slices

import (
	"fmt"

	"golang.org/x/exp/slices"
)

/*
	スライスがソート済みか判定する(IsSorted)

	func IsSorted(x []E) bool
	  param:
	    x: スライス
	  return:
	    bool: 判定結果
*/
func SlicesSample014() {
	fmt.Println("slices_sample_014")

	s1 := []int{1, 2, 3}
	s2 := []int{3, 2, 1}
	fmt.Println(slices.IsSorted(s1))
	fmt.Println(slices.IsSorted(s2))
}

/*
  実行結果
  -------
  true
  false
  -------
*/
