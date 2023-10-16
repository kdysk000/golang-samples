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

	s := make([]int, 0, 10)
	fmt.Printf("cap: %d, len: %d\n", cap(s), len(s))
	
	s = append(s, 1)
	fmt.Printf("cap: %d, len: %d\n", cap(s), len(s))
	
	s = slices.Clip(s)
	fmt.Printf("cap: %d, len: %d\n", cap(s), len(s))
}

/*
  実行結果
  -------
  3 true
  -------
*/
