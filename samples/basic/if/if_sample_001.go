package myif

import (
	"fmt"
)

/*
	基本的なif文
*/
func IfSample001() {
	fmt.Println("if_sample_001")

	num := 2

	if num == 0 {
		fmt.Println("num is 0")
	} else if num == 1 {
		fmt.Println("num is 1")
	} else {
		fmt.Println("num is 2")
	}
}

/*
  実行結果
  -------
  num is 2
  -------
*/
