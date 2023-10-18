package strings

import (
	"fmt"
	"strings"
)

/*
	文字列の分割 分割数指定版(SplitN)

	SplitN(s, sep string, n int) []string
	  概要
	    sをsepで分割したスライスを返す
	    nは分割数を指定する
	    nに負の数を指定した場合はSplitと同義
	  param:
	    s       : 文字列
	    sep     : 区切り文字列
	    n       : 分割数
	  return:
	    []string: 分割した文字列のスライス
*/
func StringsSample022() {
	fmt.Println("strings_sample_022")

	str := "hoge,fuga,foo,bar"

	str1 := strings.SplitN(str, ",", 2)
	str2 := strings.SplitN(str, "", -1)

	fmt.Println(str1)
	fmt.Println(str2)
}

/*
  実行結果
  -------
  [hoge fuga,foo,bar]
  [h o g e , f u g a , f o o , b a r]
  -------
*/
