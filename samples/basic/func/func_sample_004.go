package myfunc

import (
	"fmt"
)

/*
	戻り値を破棄する

	関数からの戻り値が不要な場合は、「 _ 」を使うことで戻り値を破棄することができる
*/
func FuncSample004() {
	fmt.Println("func_sample_004")

	//減算結果は破棄
	add1, _ := addMinus(5, 2)
	fmt.Println(add1)
}

/*
  実行結果
  -------
  7
  -------
*/
