package strings

import (
	"fmt"
	"strings"
)

/*
	文字列の置換(Replace)

	Replace(s, old, new string, n int) string
	  概要
	    sの中でoldと一致する文字列をnewに置換する
	    nに負の値を指定した場合は全て置換
	  param:
	    s     : 文字列
	    old   : 置換前の文字列
	    new   : 置換後の文字列
	    n     : 置換回数
	  return:
	    string: 置換後の文字列
*/
func StringsSample018() {
	fmt.Println("strings_sample_018")

	str := "hoge hoge hoge"

	str1 := strings.Replace(str, "hoge", "fuga", 1)
	str2 := strings.Replace(str, "hoge", "fuga", 2)
	str3 := strings.Replace(str, "hoge", "fuga", -1)

	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println(str3)
}

/*
  実行結果
  -------
  fuga hoge hoge
  fuga fuga hoge
  fuga fuga fuga
  -------
*/
