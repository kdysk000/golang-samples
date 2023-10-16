package slice

import (
	"fmt"
)

/*
	スライスの要素を削除する

	append()を使う
	  s = append(s[:i], s[i+1:]...)

	appendは第一引数に与えたスライスに、第二引数以降の値を追加する関数になるので、
	ここで行っている操作はi -1 番目までのスライスに、i+1番目以降の値を追加する。(第一引数で指定したスライスに対して処理される)
*/
func SliceSample004() {
	fmt.Println("slice_sample_004")

	s1 := []int{0,1,2,3,4,5,6,7,8,9}

	//最初の2つの要素を削除
	s2 := s1[2:]
	fmt.Println(s2)

	//最後の要素を削除
	s3 := s1[:len(s1)-1]
	fmt.Println(s3)

	//5を削除(append)
	s4 := append(s1[:5], s1[6:]...)
	fmt.Println(s4)
	fmt.Println(s1)  //s1が変化していることを確認
}

/*
  実行結果
  -------
  [2 3 4 5 6 7 8 9]
  [0 1 2 3 4 5 6 7 8]
  [0 1 2 3 4 6 7 8 9]
  [0 1 2 3 4 6 7 8 9 9]
  -------
*/
