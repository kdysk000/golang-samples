package slices

import (
	"fmt"

	"golang.org/x/exp/slices"
)

/*
	スライスの要素を削除する(Compact)

	func Delete(s S, i, j int) S
		param:
		  s: スライス
		  i: 削除する要素の先頭index
		  j: 削除する要素の終端index
		return:
		  S: 削除後の新しいスライス

	注：引数に渡す元のスライスに影響する
*/
func SlicesSample009() {
	fmt.Println("slices_sample_009")

	s1 := []int{1, 2, 3, 4, 5}
	s2 := slices.Delete(s1, 0, 1)
	fmt.Println(s1)
	fmt.Println(s2)
}

/*
  実行結果
  -------
  [2 3 4 5 5]
  [2 3 4 5]
  -------
*/
