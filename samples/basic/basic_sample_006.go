package basic

import (
	"fmt"
)

/*
	mapの宣言と初期化

	mapとは
	　・データ構造の一つで、キーと値を保持する
	　・他の言語で「連想配列」「ディクショナリ」などと呼ばれるもの

	書式
	  変数名 := map[キーのデータ型]値のデータ型{}
*/
func BasicSample006() {
	fmt.Println("basic_sample_006")

	// mapの定義と初期化
	m := map[string]int{
		"apple" : 1,
		"orange": 2,
		"banana": 3,
	}

	fmt.Println(m)
}

/*
  実行結果
  -------
  [ABC DEF GHI]
  -------
*/
