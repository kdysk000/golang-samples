package slices

import (
	"fmt"

	"golang.org/x/exp/slices"
)

/*
	スライスに要素が含まれているか判定(Contains)

	func Contains(s []E, v E) bool
		param:
		  s: スライス
		  v: 検索する要素
		return:
		  bool: みつかったかどうか
*/
func SlicesSample008() {
	fmt.Println("slices_sample_008")

	s := []int{1, 2, 3, 4, 5}
	target := 4
	ret := slices.Contains(s, target)
	fmt.Println(ret)
}

/*
  実行結果
  -------
  true
  -------
*/
