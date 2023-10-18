package strings

import (
	"fmt"
	"strings"
)

/*
	文字列中で検索文字列が最後に見つかる位置を返す(LastIndex)

	LastIndex(s, substr string) int
	  概要
	    sの中でsubstrが最後に見つかる位置を返す
	  param:
	    s     : 文字列
	    substr: 検索したい文字列
	  return:
	    int   : 位置
	注：
	  大文字小文字は区別される
	  charsの中のUnicodeコードポイントのどれもがsに含まれていない場合、-1を返す
*/
func StringsSample014() {
	fmt.Println("strings_sample_014")

	fmt.Println(strings.LastIndex("hogehoge", "h"))
	fmt.Println(strings.LastIndex("hogehoge", "H"))
	fmt.Println(strings.LastIndex("hogehoge", "hoge"))
	fmt.Println(strings.LastIndex("hogehoge", "ABC"))
	fmt.Println(strings.LastIndex("hogehoge", ""))
}

/*
  実行結果
  -------
  4
  -1
  4
  -1
  8
  -------
*/
