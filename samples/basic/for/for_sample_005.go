package myfor

import (
	"fmt"
)

/*
	ループの先頭に戻る(continue)

*/
func ForSample005() {
	fmt.Println("for_sample_005")

	for i := 0; i < 5; i++ {
		// i==2 の時だけそのあとのPrintlnが実行されずループ先頭に戻る
		if i == 2 {
			continue
		}
		fmt.Println(i)
	}
}

/*
  実行結果
  -------
  0
  1
  3
  4
  -------
*/
