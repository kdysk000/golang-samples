package slices

import (
	"fmt"

	"golang.org/x/exp/slices"
)

/*
	２つのスライスが等しいか比較(Equal)

	func Equal(s1, s2 []E) bool
	  param:
	    s1: スライス
	    s2: スライス
	  return:
	    bool: 等しいかどうか
*/
func SlicesSample010() {
	fmt.Println("slices_sample_010")

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	s3 := []int{1, 2, 2}
	fmt.Println(slices.Equal(s1, s2))
	fmt.Println(slices.Equal(s1, s3))
}

/*
  実行結果
  -------
  true
  false
  -------
*/
