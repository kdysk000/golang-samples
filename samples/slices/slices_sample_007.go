package slices

import (
	"fmt"

	"golang.org/x/exp/slices"
)

/*
	スライスの比較(Compare)

	func Compare(s1, s2 []E) int
	  param:
	    s1: スライス
	    s2: スライス
	  return:
	    int: 比較結果(0：同じ、1：s1>s2、-1：s1<s2)
*/
func SlicesSample007() {
	fmt.Println("slices_sample_007")

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	s3 := []int{1, 2, 2}
	fmt.Println(slices.Compare(s1, s2))
	fmt.Println(slices.Compare(s1, s3))
}

/*
  実行結果
  -------
  0
  1
  -------
*/
