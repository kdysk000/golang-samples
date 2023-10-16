package array

import (
	"fmt"
)

/*
	2次元配列の宣言と初期化、代入
*/
func ArraySample003() {
	fmt.Println("array_sample_003")

	// 空の配列を定義して後から代入
	ar1 := [2][3]string{}
	ar1[0][0] = "aaa"
	ar1[0][1] = "bbb"
	ar1[0][2] = "ccc"
	ar1[1][0] = "ddd"
	ar1[1][1] = "eee"
	ar1[1][2] = "fff"

	// 配列を定義する際に初期化
	ar2 := [2][3]string{{"AAA", "BBB", "CCC"}, {"DDD", "EEE", "FFF"}}

	fmt.Println(ar1)
	fmt.Println(ar2)
}

/*
  実行結果
  -------
  [[aaa bbb ccc] [ddd eee fff]]
  [[AAA BBB CCC] [DDD EEE FFF]]
  -------
*/
