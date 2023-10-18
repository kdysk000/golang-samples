package strings

import (
	"fmt"
	"strings"
)

/*
	文字列をセパレータのあとで分割(SplitAfter)

	SplitAfter(s, sep string) []string
	  概要
	    sをsepの後で分割したスライスを返す
	    nに負の値を指定した場合は全て分割
	  param:
	    s       : 文字列
	    sep     : 区切り文字列
	  return:
	    []string: 分割した文字列のスライス
*/
func StringsSample021() {
	fmt.Println("strings_sample_021")

	str := "hoge,fuga,foo,bar"

	str1 := strings.SplitAfter(str, ",")
	str2 := strings.SplitAfter(str, "")

	fmt.Println(str1)
	fmt.Println(str2)
}

/*
  実行結果
  -------
  [hoge, fuga, foo, bar]
  [h o g e , f u g a , f o o , b a r]
  -------
*/
