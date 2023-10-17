package slices

import (
	"fmt"

	"golang.org/x/exp/slices"
)

/*
	スライス要素の検索(BinarySearchFunc)

	func BinarySearchFunc(x []E, target T, cmp func(E, T) int) (int, bool)
	  param:
	    x     : 検索対象のスライス
	    target: 検索する要素
	    cmp   : 比較関数
	  return:
	    int   : ターゲットの位置
	    bool  : みつかったかどうか

	BinarySearchの比較部分を指定できる

*/
func SlicesSample002() {
	fmt.Println("slices_sample_002")

	s := []int{1, 2, 3, 4, 5}
	target := 4
	i, found := slices.BinarySearchFunc(s, target, func(x, target int) int {
		if x < target {
			return -1
		} else if x == target {
			return 0
		} else {
			return 1
		}
	})
	fmt.Println(i, found)
}

/*
  実行結果
  -------
  3 true
  -------
*/
