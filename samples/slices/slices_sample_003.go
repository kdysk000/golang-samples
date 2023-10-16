package slices

import (
	"fmt"

	"golang.org/x/exp/slices"
)

/*
	スライスの空いているCapacity を削除する(Clip)

	func Clip(s S) S
		param:
		  s: スライス
		return:
		  S: 処理後のスライス
*/
func SlicesSample003() {
	fmt.Println("slices_sample_003")

	x := make([]int, 0, 10)
	fmt.Printf("cap: %d, len: %d\n", cap(x), len(x))
	
	x = append(x, 1)
	fmt.Printf("cap: %d, len: %d\n", cap(x), len(x))
	
	x = slices.Clip(x)
	fmt.Printf("cap: %d, len: %d\n", cap(x), len(x))
}

/*
  実行結果
  -------
  3 true
  -------
*/
