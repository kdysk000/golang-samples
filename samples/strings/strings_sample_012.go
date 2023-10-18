package strings

import (
	"fmt"
	"strings"
)

/*
	文字列中で検索文字列が最初に見つかる位置を返す(IndexAny)

	IndexAny(s, chars string) int
	  概要
	    sの中でcharsに含まれるUnicodeコードポイントのいずれかが最初に見つかる位置を返す
	  param:
	    s    : 文字列
	    chars: 検索したいUnicodeコードポイントの文字列
	  return:
	    int  : 位置
	注：
	  大文字小文字は区別される
	  charsの中のUnicodeコードポイントのどれもがsに含まれていない場合、-1を返す
*/
func StringsSample012() {
	fmt.Println("strings_sample_012")

	fmt.Println(strings.IndexAny("hogefuga", "h"))
	fmt.Println(strings.IndexAny("hogefuga", "H"))
	fmt.Println(strings.IndexAny("hogefuga", "ge"))
	fmt.Println(strings.IndexAny("hogefuga", "ABC"))
	fmt.Println(strings.IndexAny("hogefuga", ""))
}

/*
  実行結果
  -------
  0
  -1
  2
  -1
  -1
  -------
*/
