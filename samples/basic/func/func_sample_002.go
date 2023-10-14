package myfunc

import (
	"fmt"
)

/*
	複数戻り値の関数の定義

	複数の値を返却することもできる。複数の場合は型名は (...) で囲む

	書式：
	    func 関数名 (引数 型) (型, 型) {
			// do something
			return 値
		}
*/
func FuncSample002() {
	fmt.Println("func_sample_002")

	add1, minus1 := addMinus(5, 2)
	fmt.Println(add1, minus1)

	// 不要な戻り値の場合は _ を使用する
	add2, _ := addMinus(5, 3)
	fmt.Println(add2)
}

func addMinus(x int, y int) (int, int) {
	return x + y, x - y
}

/*
  実行結果
  -------
  7 3
  8
  -------
*/
