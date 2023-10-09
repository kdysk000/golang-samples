package basic

import (
	"fmt"
)

/*
	可変引数の関数の定義

	... を用いることで可変引数にすることができる
*/
func BasicSample019() {
	fmt.Println("basic_sample_019")

	fmt.Println(add(1, 2))
	fmt.Println(add(1, 2, 3))
	fmt.Println(add(1, 2, 3, 4))
}

func add(x int, y ... int) int {
    for _, num := range y {
		x += num
    }
	return x
}

/*
  実行結果
  -------
  3
  6
  10
  -------
*/
