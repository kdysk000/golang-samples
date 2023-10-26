package array

import (
	"fmt"
)

/*
	インデックスを指定して初期化

	初期化しなかったインデックスの要素は初期値となる

	書式
	  ar := [...]int{1:1, 3:3, 5:5}
*/
func ArraySample003() {
	fmt.Println("array_sample_003")

	// インデックス指定あり
	ar := [...]int{1:1, 3:3, 5:5}

	fmt.Println("ar:", ar)
}

/*
  実行結果
  -------
  ar: [0 1 0 3 0 5]
  -------
*/
