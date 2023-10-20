package myif

import (
	"fmt"
)

/*
	簡易文付きif文

	書式：
	  if 変数名 := 値; 条件式 {
	  	// trueなら処理する
	  }

	注：
	  簡易付きif文の簡易文に設定している変数は、簡易文付きif文の中だけで使用可能
*/
func IfSample002() {
	fmt.Println("if_sample_002")

	num := 5

	if ret := mul(num); ret == 1 {
		fmt.Println(num, "is odd number")
	} else {
		fmt.Println(num, "is even number")
	}
}

func mul(num int) int {
	return num % 2
}

/*
  実行結果
  -------
  5 is odd number.
  -------
*/
