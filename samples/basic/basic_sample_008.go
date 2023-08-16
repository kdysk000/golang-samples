package basic

import (
	"fmt"
)

/*
	mapの要素を削除する(delete)
*/
func BasicSample008() {
	fmt.Println("basic_sample_008")

	m := map[string]int{
		"apple" : 1,
		"orange": 2,
		"banana": 3,
	}

	fmt.Println(m)

	delete(m, "banana")

	fmt.Println(m)
}

/*
  実行結果
  -------
  map[apple:1 banana:3 orange:2]
  map[apple:1 orange:2]
  -------
*/
