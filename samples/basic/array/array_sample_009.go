package array

import (
	"fmt"
)

/*
	配列のコピー(copy)

	書式
	  func copy(dst, src []Type) int
	    param:
	      dst: コピー先配列
		  src: コピー元配列
	    return:
	      int: コピーした要素数

	ただし、copyの2つの引数(dst,src)はスライスを渡す必要があるため、配列をスライス式([:])で引数に渡す
*/
func ArraySample009() {
	fmt.Println("array_sample_009")

	ar1 := [5]int{1,2,3,4,5}
	ar2 := [5]int{}

	ret := copy(ar2[:], ar1[:])
	fmt.Println("ar2:", ar2)
	fmt.Println("ret:", ret)
}

/*
  実行結果
  -------
  ar2: [1 2 3 4 5]
  ret: 5
  -------
*/
