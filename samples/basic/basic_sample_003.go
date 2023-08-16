package basic

import (
	"fmt"
)

/*
	スライスの宣言と初期化

	スライスとは
	　・配列とは違い個数をあとから変更可能で動的な配列として振る舞う

	・型名の前に [] をつけて宣言する
	・配列と同じくインデックスは 0 から始まる
	・スライスは == で比較できない
*/
func BasicSample003() {
	fmt.Println("basic_sample_003")

	// スライスを定義する際に初期化
	s1 := []string{"ABC", "DEF", "GHI"}

	fmt.Println(s1)
}

/*
  実行結果
  -------
  [ABC DEF GHI]
  -------
*/
