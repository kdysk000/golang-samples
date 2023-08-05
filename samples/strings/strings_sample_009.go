package strings

import (
	"fmt"
	"strings"
)

/*
	文字列がprefixで始まるかどうかを判定(HasPrefix)
	HasPrefix(s, prefix string) bool
		param:
		  s     : 文字列
		  prefix: 判定したいプレフィックス
		return:
		  bool  :
	注：
	  大文字小文字は区別される
	  prefixに空文字が指定されていたらtrueが返る
*/
func StringsSample009() {
	fmt.Println("strings_sample_009")

	fmt.Println(strings.HasPrefix("hogehoge", "h"))
	fmt.Println(strings.HasPrefix("hogehoge", "H"))
	fmt.Println(strings.HasPrefix("hogehoge", "hoge"))
	fmt.Println(strings.HasPrefix("hogehoge", "fuga"))
	fmt.Println(strings.HasPrefix("hogehoge", ""))
}

/*
  実行結果
  -------
  true
  false
  true
  false
  true
  -------
*/
