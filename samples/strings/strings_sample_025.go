package strings

import (
	"fmt"
	"strings"
)

/*
	文字列の先頭から不要なUnicode コードポイントを除去(TrimLeft)
	TrimLeft(s, cutset string) string
		概要
		  sの先頭からcutsetに含まれるUnicodeコードポイントを除いた文字列を返す
		param:
		  s     : 文字列
		  cutset: 除去するUnicodeコードポイントを含んだ文字列
		return:
		  string: 除去後の文字列
*/
func StringsSample025() {
	fmt.Println("strings_sample_025")

	str := "@! Hello world. @@@!! "

	str1 := strings.TrimLeft(str, "@! ")

	fmt.Println(str1)
}

/*
  実行結果
  -------
  Hello world. @@@!! 
  -------
*/
