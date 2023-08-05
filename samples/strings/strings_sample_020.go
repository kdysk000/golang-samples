package strings

import (
	"fmt"
	"strings"
)

/*
	文字列の分割(Split)
	Split(s, sep string) []string
		概要
		  sをsepで分割したスライスを返す
		  sの中にsepが含まれていない場合はsのみを含む長さが1のスライスを返す
		  sepが空文字の場合はsをUTF-8文字ごとに分割したスライスを返す
		param:
		  s       : 文字列
		  sep     : 区切り文字列
		return:
		  []string: 分割した文字列のスライス
*/
func StringsSample020() {
	fmt.Println("strings_sample_020")

	str := "hoge,fuga,foo,bar"

	str1 := strings.Split(str, ",")
	str2 := strings.Split(str, "")

	fmt.Println(str1)
	fmt.Println(str2)
}

/*
  実行結果
  -------
  [hoge fuga foo bar]
  [h o g e , f u g a , f o o , b a r]
  -------
*/
