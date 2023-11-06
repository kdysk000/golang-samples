package myfunc

import (
	"fmt"
)

/*
	クロージャ

	関数の中で作成される関数として実装されて、関数作成時点の変数を閉じ込めることができる。

	下記の例だと、counter()を実行するごとに i がインクリメントされその値が返る
*/
func FuncSample013() {
	fmt.Println("func_sample_013")

	counter := func() func() int {
		i := 0
		return func() int{
			i++
			return i
		}
	}()

	fmt.Println(counter())
	fmt.Println(counter())
}

/*
  実行結果
  -------
  1
  2
  -------
*/
