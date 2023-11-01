package array

import (
	"fmt"
)

/*
	配列の一部を部分配列として取り出す

	書式
	  a[開始index:終了index]

	例
	　a[0:2]  index0～index1
	  a[1:]   index1～最後まで
	  a[:5]   index0～index4
*/
func ArraySample008() {
	fmt.Println("array_sample_008")

	ar := [10]int{1,2,3,4,5,6,7,8,9,10}

	fmt.Println(ar[:2])          //先頭から2つ
	fmt.Println(ar[len(ar)-2:])  //末尾2つ
	fmt.Println(ar[2:6])         //index2～5
}

/*
  実行結果
  -------
  [1 2]
  [9 10]
  [3 4 5 6]
  -------
*/
