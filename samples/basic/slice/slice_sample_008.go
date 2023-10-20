package slice

import (
	"fmt"
)

/*
	スライスの一部を部分スライスとして取り出す

	書式
	  a[開始index:終了index]

	例
	　a[0:2]  index0～index1
	  a[1:]   index1～最後まで
	  a[:5]   最初～index4
*/
func SliceSample008() {
	fmt.Println("slice_sample_008")

	s := []int{0,1,2,3,4,5,6,7,8,9}

	fmt.Println(s[:2])         //先頭から2つ
	fmt.Println(s[len(s)-2:])  //末尾2つ

	fmt.Println(s[2:6])        //index2～5
}

/*
  実行結果
  -------
  [0 1]
  [8 9]
  [2 3 4 5]
  -------
*/
