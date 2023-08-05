package strings

import (
	"fmt"
	"strings"
)

/*
	文字列中で検索文字が最後に見つかる位置を返す(LastIndexByte)
	LastIndexByte(s string, c byte) int
		概要
		  sの中でcが最後に見つかる位置を返す
		param:
		  s  : 文字列
		  c  : 検索したい文字
		return:
		  int: 位置
	注：
	  大文字小文字は区別される
	  cがsに含まれていない場合、-1を返す
*/
func StringsSample016() {
	fmt.Println("strings_sample_016")

	fmt.Println(strings.LastIndexByte("hogehoge", 'h'))
	fmt.Println(strings.LastIndexByte("hogehoge", 'H'))
	fmt.Println(strings.LastIndexByte("hogehoge", 'e'))
}

/*
  実行結果
  -------
  4
  -1
  7
  -------
*/
