package myfor

import (
	"fmt"
)

/*
	while文に似た書き方のfor文

	Go言語にはwhile文がないが、while文に似た書き方でfor文を書くことができる

	書式：
	  初期値
	  for 条件 {
	  	// 繰り返し処理する内容
	  	// 増減
	  }
*/
func ForSample002() {
	fmt.Println("for_sample_002")

    i := 0
    for i < 5 {
        fmt.Println(i)
        i++
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
