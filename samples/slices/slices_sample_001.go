package slices

import (
	"fmt"

	"golang.org/x/exp/slices"
)

/*
	スライス要素の検索(BinarySearch)

	func BinarySearch(x []E, target E) (int, bool)
		param:
		  x     : 検索対象のスライス
		  target: 検索する要素
		return:
		  int : ターゲットの位置
		  bool: みつかったかどうか

	スライスのターゲットの位置と見つかったかを返す

*/
func SlicesSample001() {
	fmt.Println("slices_sample_001")

	s := []int{1, 2, 3, 4, 5}
	target := 4
	i, found := slices.BinarySearch(s, target)
	fmt.Println(i, found)
}

/*
  実行結果
  -------
  3 true
  -------
*/
