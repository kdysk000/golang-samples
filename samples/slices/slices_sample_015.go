package slices

import (
	"fmt"

	"golang.org/x/exp/slices"
)

/*
	スライスの要素を置換する(Replace)

	func Replace(s S, i, j int, v …E) S
	  param:
	    s: スライス
	    i: 置換する要素の先頭index
	    j: 置換する要素の終端index
	    v: 要素
	  return:
	    S: 置換後のスライス

	※いまいち仕様がわからないので使えないかも・・・
*/
func SlicesSample015() {
	fmt.Println("slices_sample_015")

	s1 := []int{1, 2, 3}
	s2 := slices.Replace(s1, 0, 1, 9, 9) //Index0,1を9に置換したいけど・・
	fmt.Println(s1)
	fmt.Println(s2)  //0だけ置換されてそのあとに9が挿入されてる？

	s3 := slices.Replace(s1, 0, 1, 9)
	fmt.Println(s1)  //元のスライスに影響している
	fmt.Println(s3)  //0だけ置換されて2はそのまま
}

/*
  実行結果
  -------
  [1 2 3]
  [9 9 2 3]
  [9 2 3]
  [9 2 3]
  -------
*/
