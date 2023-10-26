package array

import (
	"fmt"
)

/*
	配列の特定の要素を取り出す
*/
func ArraySample007() {
	fmt.Println("array_sample_007")

	ar := [5]int{1,2,3,4,5}

	fmt.Println(ar[0])           //先頭
	fmt.Println(ar[len(ar)-1])   //末尾
}

/*
  実行結果
  -------
  1
  5
  -------
*/
