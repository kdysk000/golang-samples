package basic

import (
	"fmt"
)

/*
	スライスに要素を追加する(append)

	s2 = append(s1, 追加する値1, 追加する値2, ・・・)
	  s1: 値を追加したいスライス
	  s2: 追加後のスライス

	あるスライスに別のスライスの内容を全て追加するには ... を用いる
*/
func BasicSample004() {
	fmt.Println("basic_sample_004")

	// 空のスライスを定義したあとappendで追加
	s1 := []string{}
	s1 = append(s1, "ABC")

	s2 := []string{"DEF", "GHI"}

	// 2つのスライスを結合
	s3 := append(s1, s2...)

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}

/*
  実行結果
  -------
  [ABC]
  [DEF GHI]
  [ABC DEF GHI]
  -------
*/
