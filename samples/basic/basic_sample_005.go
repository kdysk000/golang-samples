package basic

import (
	"fmt"
)

/*
	make()を用いたスライスの作成

	書式
	  変数名 := make([]データ型, 長さ, キャパシティ)

	・型と長さは必須、キャパシティは任意(キャパシティを省略した場合は長さと同じ値になる)
	・長さで指定した箇所はそれぞれのデータ型の既定値で初期化される

*/
func BasicSample005() {
	fmt.Println("basic_sample_005")

	// スライスの作成
	s1 := make([]int, 5)
	s2 := make([]int, 3, 5)
	s3 := make([]int, 0, 5)

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}

/*
  実行結果
  -------
  [0 0 0 0 0]
  [0 0 0]
  []
  -------
*/
