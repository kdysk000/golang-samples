package slices

import (
	"fmt"

	"golang.org/x/exp/slices"
)

/*
	スライスのCapacityを増やす(Grow)

	func Grow(s S, n int) S
		param:
		  s: スライス
		  n: 増やした後のCapacity数
		return:
		  S: 新しいスライス
*/
func SlicesSample011() {
	fmt.Println("slices_sample_011")

	s1 := make([]int, 0, 1)
	s2 := slices.Grow(s1, 3)
	fmt.Println(cap(s1))
	fmt.Println(cap(s2))
}

/*
  実行結果
  -------
  1
  3
  -------
*/
