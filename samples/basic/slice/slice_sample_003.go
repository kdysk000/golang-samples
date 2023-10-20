package slice

import (
	"fmt"
)

/*
	スライスに要素を追加する(append)

	書式
	  s2 = append(s1, 追加する値1, 追加する値2, ・・・)
	    s1: 値を追加したいスライス
	    s2: 追加後のスライス
*/
func SliceSample003() {
	fmt.Println("slice_sample_003")

	// 空のスライスを定義したあとappendで追加
	s := []string{}
	s = append(s, "ABC")

	fmt.Println(s)
}

/*
  実行結果
  -------
  [ABC]
  -------
*/
