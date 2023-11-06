package strings

import (
	"fmt"
	"strings"
)

/*
	文字列中で検索文字列が最後に見つかる位置を返す(LastIndexAny)

	LastIndexAny(s, chars string) int
	  概要
	    sの中でcharsに含まれるUnicodeコードポイントのいずれかが最後に見つかる位置を返す
	  param:
	    s    : 文字列
	    chars: 検索したいUnicodeコードポイントの文字列
	  return:
	    int  : 位置
	注：
	  大文字小文字は区別される
	  charsの中のUnicodeコードポイントのどれもがsに含まれていない場合、-1を返す
*/
func StringsSample015() {
	fmt.Println("strings_sample_015")

	fmt.Println(strings.LastIndexAny("hogehoge", "h"))
	fmt.Println(strings.LastIndexAny("hogehoge", "H"))
	fmt.Println(strings.LastIndexAny("hogehoge", "ge"))
	fmt.Println(strings.LastIndexAny("hogehoge", "ABC"))
	fmt.Println(strings.LastIndexAny("hogehoge", ""))
}

/*
  実行結果
  -------
  4
  -1
  7
  -1
  -1
  -------
*/
