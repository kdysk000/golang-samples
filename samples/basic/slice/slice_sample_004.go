package slice

import (
	"fmt"
)

/*
	スライスを結合する(append)

	書式
	  s3 := append(s1, s2...)
	    s1：スライス1
	    s2：スライス2
	    s3：結合後のスライス

	あるスライスに別のスライスの内容を全て追加するには ... を用いる
*/
func SliceSample004() {
	fmt.Println("slice_sample_004")

	s1 := []int{1,2,3,4,5}
	s2 := []int{6,7,8,9,10}

	s3 := append(s1, s2...)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}

/*
  実行結果
  -------
  [1 2 3 4 5]
  [6 7 8 9 10]
  [1 2 3 4 5 6 7 8 9 10]
  -------
*/
