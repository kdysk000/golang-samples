package slices

import (
	"fmt"

	"golang.org/x/exp/slices"
)

/*
	スライスに指定したindexへ要素を追加する(Insert)

	func Insert(s S, i int, v …E) S
		param:
		  s: スライス
		  i: 追加するIndex
		  v: 追加する要素
		return:
		  S: 追加後のスライス
*/
func SlicesSample013() {
	fmt.Println("slices_sample_013")

	s1 := []int{1, 2, 5}
	s2 := slices.Insert(s1, 2, 3, 4)
	fmt.Println(s1)
	fmt.Println(s2)
}

/*
  実行結果
  -------
  [1 2 5]
  [1 2 3 4 5]
  -------
*/
