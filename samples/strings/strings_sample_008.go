package strings

import (
	"fmt"
	"strings"
)

/*
	文字列を1つ以上のfを満たすUnicodeコードポイントで分割した文字列のスライスを返す(FieldsFunc)

	FieldsFunc(s string, f func(rune) bool) []string
	  param:
	    s       : 文字列
	    f       : 判定関数
	  return:
	    []string: 文字列のスライス

	注：
	  すべてのコードポイントがfを満たす、もしくはsが空文字の場合は空のスライスを返す
*/
func StringsSample008() {
	fmt.Println("strings_sample_008")

	f := func(r rune) bool {
		return r == '#'
	}

	str1 := strings.FieldsFunc("###hoge#fuga##foo##bar##", f)
	str2 := strings.FieldsFunc("####", f)
	str3 := strings.FieldsFunc("", f)

	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println(str3)

}

/*
  実行結果
  -------
  [hoge fuga foo bar]
  []
  []
  -------
*/
