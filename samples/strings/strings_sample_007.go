package strings

import (
	"fmt"
	"strings"
)

/*
	文字列を1つ以上の空白で分割した文字列のスライスを返す(Fields)
	Fields(s string) []string
		param:
		  s       : 文字列
		return:
		  []string: 文字列のスライス
*/
func StringsSample007() {
	fmt.Println("strings_sample_007")

	str := strings.Fields("   hoge fuga\nfoo\tbar\rpiyo   ")
	fmt.Println(str)
}

/*
  実行結果
  -------
  [hoge fuga foo bar piyo]
  -------
*/
