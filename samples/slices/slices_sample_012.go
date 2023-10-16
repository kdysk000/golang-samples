package slices

import (
	"fmt"

	"golang.org/x/exp/slices"
)

/*
	一致するスライスの要素のIndexを返す(Index)

	func Index(s []E, v E) int
		param:
		  s: スライス
		  v: 検索する要素
		return:
		  int: Index(見つかった：Index番号、みつからなかった：-1)
*/
func SlicesSample012() {
	fmt.Println("slices_sample_012")

	s := []int{1, 2, 3, 4, 5}
	target := 4
	fmt.Println(slices.Index(s, target))
	fmt.Println(slices.Index(s, 0))
}

/*
  実行結果
  -------
  3
  -1
  -------
*/
