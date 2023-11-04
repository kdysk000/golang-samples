package myfunc

import (
	"fmt"
)

/*
	値渡しと参照渡し

	Goでは数値型(int,float64など)、文字列型、構造体などの型では基本的に値渡しとなる。
	値渡しでは、関数を呼び出す時の引数には値がコピーされ、関数内ではそのコピーされた値が使われるため、
	関数内で値を書き換えても元の変数には影響を与えない。

	一方、マップ(map)やスライス(slice)の場合は参照渡しになるため、
	関数内で値を書き換えると元の変数が変更される。
*/
func FuncSample012() {
	fmt.Println("func_sample_012")

	a := 1
	b := 1.1
	c := []int{0,1,2}

	testFunc4(a, b, c)

	fmt.Println(a)  // aは値渡しなので変更されていない
	fmt.Println(b)  // bは値渡しなので変更されていない
	fmt.Println(c)  // cは参照渡しなので変更されている
}

func testFunc4(a int, b float64, c []int)  {
	a = 10
	b = 10.1
	c[0] = 10
}

/*
  実行結果
  -------
  1
  1.1
  [10 1 2]
  -------
*/
