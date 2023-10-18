package strings

import (
	"fmt"
	"strings"
)

/*
	文字列がsuffixで終わるかどうかを判定(HasSuffix)

	HasSuffix(s, suffix string) bool
	  param:
	    s     : 文字列
	    suffix: 判定したいサフィックス
	  return:
	    bool  :
	注：
	  大文字小文字は区別される
	  suffixに空文字が指定されていたらtrueが返る
*/
func StringsSample010() {
	fmt.Println("strings_sample_010")

	fmt.Println(strings.HasSuffix("hogehoge", "e"))
	fmt.Println(strings.HasSuffix("hogehoge", "E"))
	fmt.Println(strings.HasSuffix("hogehoge", "hoge"))
	fmt.Println(strings.HasSuffix("hogehoge", "fuga"))
	fmt.Println(strings.HasSuffix("hogehoge", ""))
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
