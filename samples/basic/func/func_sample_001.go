package myfunc

import (
	"fmt"
)

/*
	関数の定義(func)

	書式：
	  func 関数名 (引数 型) 型 {
	  	// do something
	  	return 値
	  }

	Public関数とPrivate関数
	  Public関数：関数名が大文字で始まる。他のパッケージから当該関数をコールすることが可能
	  Private関数：関数名が小文字で始まる。定義したパッケージからのみコールすることが可能
}
*/
func FuncSample001() {
	fmt.Println("func_sample_001")
	fmt.Println(minus(5, 2))
}

func minus(x int, y int) int {
	return x - y
}

/*
  実行結果
  -------
  3
  -------
*/
