package basic

import (
	"fmt"
)

/*
	遅延実行(defer)

	関数が終了される直前実行されることが保証される関数
*/
func BasicSample020() {
	fmt.Println("basic_sample_020")

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
