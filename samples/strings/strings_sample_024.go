package strings

import (
	"fmt"
	"strings"
)

/*
	文字列の先頭と末尾から不要なUnicode コードポイントを除去(Trim)
	Trim(s, cutset string) string
		概要
		  sの先頭と末尾からcutsetに含まれるUnicodeコードポイントを除いた文字列を返す
		param:
		  s     : 文字列
		  cutset: 除去するUnicodeコードポイントを含んだ文字列
		return:
		  string: 除去後の文字列
*/
func StringsSample024() {
	fmt.Println("strings_sample_024")

	str := "@! Hello world. @@@!! "

	str1 := strings.Trim(str, "@! ")
	str2 := strings.Trim(str, "")

	fmt.Println(str1)
	fmt.Println(str2)
}

/*
  実行結果
  -------
  Hello world.
  @! Hello world. @@@!! 
  -------
*/
