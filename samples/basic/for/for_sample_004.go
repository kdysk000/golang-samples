package myfor

import (
	"fmt"
)

/*
	ループの中断(break)

*/
func ForSample004() {
	fmt.Println("for_sample_004")

	for i := 0; i < 5; i++ {
		fmt.Println(i)
		if i == 2 {
			break
		}
	}
}

/*
  実行結果
  -------
  0
  1
  2
  -------
*/
