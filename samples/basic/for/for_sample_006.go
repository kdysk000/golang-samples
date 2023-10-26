package myfor

import (
	"fmt"
)

/*
	ループ変数が共有される問題

	・Goのfor文で宣言されるループ変数は、ループ毎に同じアドレスに値が格納される仕様になっている
	・これによりループ内で並行処理を実行するなど特定のケースで、どのループでも毎回同じ値が取得されてしまうといった意図しないバグが発生する
	・Go 1.21でこの問題への対策のプレビューが公開され、1.22で正式リリースされる予定

	詳細は以下の記事を参照
	https://moneyforward-dev.jp/entry/2023/08/15/150000
*/
func ForSample006() {
	fmt.Println("for_sample_006")

	values := []int{0, 2, 4, 6}
	for i, v := range values {
		fmt.Printf("index: %d, value: %d, address: %p\n", i, v, &v)
	}
}

/*
  実行結果
  -------
  index: 0, value: 0, address: 0xc0000180e0
  index: 1, value: 2, address: 0xc0000180e0
  index: 2, value: 4, address: 0xc0000180e0
  index: 3, value: 6, address: 0xc0000180e0
  -------
*/
