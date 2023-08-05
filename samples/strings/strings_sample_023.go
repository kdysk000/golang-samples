package strings

import (
	"fmt"
	"strings"
)

/*
	文字列をセパレータのあとで分割 分割数指定版(SplitAfterN)
	SplitAfterN(s, sep string, n int) []string
		概要
		  sをsepの後で分割したスライスを返す
		  nに負の数を指定した場合はSplitAfterと同義
		param:
		  s       : 文字列
		  sep     : 区切り文字列
		return:
		  []string: 分割した文字列のスライス
*/
func StringsSample023() {
	fmt.Println("strings_sample_023")

	str := "hoge,fuga,foo,bar"

	str1 := strings.SplitAfterN(str, ",", 2)
	str2 := strings.SplitAfterN(str, "", -1)

	fmt.Println(str1)
	fmt.Println(str2)
}

/*
  実行結果
  -------
  [hoge, fuga,foo,bar]
  [h o g e , f u g a , f o o , b a r]
  -------
*/
