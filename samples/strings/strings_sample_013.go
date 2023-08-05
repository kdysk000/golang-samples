package strings

import (
	"fmt"
	"strings"
)

/*
	文字列中で検索文字が最初に見つかる位置を返す(IndexByte)
	IndexByte(s string, c byte) int
		概要
		  sの中でcが最初に見つかる位置を返す
		param:
		  s  : 文字列
		  c  : 検索したい文字
		return:
		  int: 位置
	注：
	  大文字小文字は区別される
	  sの中にcが含まれていない場合、-1を返す
*/
func StringsSample013() {
	fmt.Println("strings_sample_013")

	fmt.Println(strings.IndexByte("hogehoge", 'h'))
	fmt.Println(strings.IndexByte("hogehoge", 'H'))
	fmt.Println(strings.IndexByte("hogehoge", 'e'))
}

/*
  実行結果
  -------
  0
  -1
  3
  -------
*/
