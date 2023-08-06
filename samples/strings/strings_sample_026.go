package strings

import (
	"fmt"
	"strings"
)

/*
	文字列の末尾から不要なUnicode コードポイントを除去(TrimRight)
	TrimRight(s, cutset string) string
		概要
		  sの末尾からcutsetに含まれるUnicodeコードポイントを除いた文字列を返す
		param:
		  s     : 文字列
		  cutset: 除去するUnicodeコードポイントを含んだ文字列
		return:
		  string: 除去後の文字列
*/
func StringsSample026() {
	fmt.Println("strings_sample_026")

	str := "@! Hello world. @@@!! "

	str1 := strings.TrimRight(str, "@! ")

	fmt.Println(str1)
}

/*
  実行結果
  -------
  @! Hello world.
  -------
*/
