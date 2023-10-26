package array

import (
	"fmt"
)

/*
	配列の初期値

	・数値は 0
	・string は ""
	・bool は false

*/
func ArraySample002() {
	fmt.Println("array_sample_002")

	// 空の配列
	ar1 := [3]int{}
	ar2 := [3]string{}
	ar3 := [3]bool{}

	fmt.Println("ar1:", ar1)
	fmt.Println("ar2:", ar2)
	fmt.Println("ar3:", ar3)
}

/*
  実行結果
  -------
  ar1: [0 0 0]
  ar2: [  ]
  ar3: [false false false]
  -------
*/
