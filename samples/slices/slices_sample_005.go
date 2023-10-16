package slices

import (
	"fmt"

	"golang.org/x/exp/slices"
)

/*
	スライスの連続した同じ要素を削除し新しいスライスを返す(Compact)

	func Compact(s S) S
		param:
		  s: スライス
		return:
		  S: 削除後の新しいスライス

	注：引数に渡す元のスライスに影響する
*/
func SlicesSample005() {
	fmt.Println("slices_sample_005")

	s1 := []int{1, 1, 1, 2, 2, 2}
	s2 := slices.Compact(s1)
	fmt.Println(s1)
	fmt.Println(s2)
}

/*
  実行結果
  -------
  [1 2]
  -------
*/
