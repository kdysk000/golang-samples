package array

import (
	"fmt"
)

/*
	配列の比較

	・== 演算子で全ての要素の値を比較し全て同じ場合にtrueが返る
	・< や > 演算子は配列の比較に使用できない
*/
func ArraySample004() {
	fmt.Println("array_sample_004")

	ar1 := [3]int{0, 1, 2}
	ar2 := [3]int{0, 1, 2}
	ar3 := [3]int{0, 1, 3}

	fmt.Println(ar1 == ar2)
	fmt.Println(ar1 == ar3)
}

/*
  実行結果
  -------
  true
  false
  -------
*/
