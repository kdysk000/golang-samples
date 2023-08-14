package strings

import (
	"fmt"
	"strings"
)

/*
	一致する全ての文字列の置換(ReplaceAll)
	ReplaceAll(s, old, new string) string
		概要
		  sの中でoldと一致する全ての文字列をnewに置換する
		param:
		  s     : 文字列
		  old   : 置換前の文字列
		  new   : 置換後の文字列
		return:
		  string: 置換後の文字列
*/
func StringsSample019() {
	fmt.Println("strings_sample_019")

	str := "hoge hoge hoge"

	str1 := strings.ReplaceAll(str, "hoge", "fuga", )

	fmt.Println(str1)
}

/*
  実行結果
  -------
  fuga fuga fuga
  -------
*/
