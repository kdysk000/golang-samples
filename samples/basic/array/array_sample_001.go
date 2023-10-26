package array

import (
	"fmt"
)

/*
	配列の宣言と初期化、代入

	・型名の前に [個数] をつけて宣言する
	・[...]とすることで個数を省略することが可能
	・配列のインデックスは 0 から始まる
	・配列はコンパイル時に個数が決まっていて変更不可
	・途中で個数を変更できないがメモリ効率や性能の点で優れているため、予め個数が決まっている場合は配列を使用する
*/
func ArraySample001() {
	fmt.Println("array_sample_001")

	// 空の配列を定義して後から代入
	ar1 := [3]string{}
	ar1[0] = "abc"
	ar1[1] = "def"
	ar1[2] = "ghi"

	// 配列を定義する際に初期化
	ar2 := [3]string{"ABC", "DEF", "GHI"}

	// 個数の省略
	ar3 := [...]string{"123", "456", "789"}

	fmt.Println("ar1:", ar1)
	fmt.Println("ar2:", ar2)
	fmt.Println("ar3:", ar3)
}

/*
  実行結果
  -------
  ar1: [abc def ghi]
  ar2: [ABC DEF GHI]
  ar3: [123 456 789]
  -------
*/
