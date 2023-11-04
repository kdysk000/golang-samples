package myfunc

import (
	"fmt"
)

/*
	遅延実行(defer)

	関数が終了される直前実行されることが保証される関数
*/
func FuncSample006() {
	fmt.Println("func_sample_006")

	defer fmt.Println("END")
	fmt.Println("START")
}

/*
  実行結果
  -------
  START
  END
  -------
*/
