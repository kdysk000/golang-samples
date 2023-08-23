package basic

import (
	"fmt"
)

/*
	基本的なfor文

	書式：
		for 初期値; 条件; 増減 {
    		// 繰り返し処理する内容
		}
*/
func BasicSample012() {
	fmt.Println("basic_sample_012")

	for i := 0; i < 5; i++ {
        fmt.Println(i)
    }
}

/*
  実行結果
  -------
  0
  1
  2
  3
  4
  -------
*/
