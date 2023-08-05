package strings

import (
	"fmt"
	"strings"
)

/*
	文字列中で検索文字列が最初に見つかる位置を返す(Index)
	Index(s, substr string) int
		param:
		  s     : 文字列
		  substr: 検索したい文字列
		return:
		  int   : 位置
	注：
	  大文字小文字は区別される
	  substrが含まれていない場合は-1を返す
*/
func StringsSample011() {
	fmt.Println("strings_sample_011")

	fmt.Println(strings.Index("hogefuga", "h"))
	fmt.Println(strings.Index("hogefuga", "H"))
	fmt.Println(strings.Index("hogefuga", "hoge"))
	fmt.Println(strings.Index("hogefuga", "fuga"))
	fmt.Println(strings.Index("hogefuga", ""))
}

/*
  実行結果
  -------
  0
  -1
  0
  4
  0
  -------
*/
