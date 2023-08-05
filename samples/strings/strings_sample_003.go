package strings

import (
	"fmt"
	"strings"
)

/*
	指定の文字列に含まれる Unicode コードポイントのいずれかが含まれているか(ContainsAny)
	func ContainsAny(s, chars string) bool
		概要:
		  例えば、charsが"ABC"なら'A'と'B'と'C'のいずれかが s に含まれるかを判定する
		param:
		  s    : 検索対象の文字列
		  chars: 検索する文字列
		return:
		  bool :
*/
func StringsSample003() {
	fmt.Println("strings_sample_003")

	fmt.Println(strings.ContainsAny("hogehoge", "e"))
	fmt.Println(strings.ContainsAny("hogehoge", "hoge"))
	fmt.Println(strings.ContainsAny("hogehoge", "HOGE"))
	fmt.Println(strings.ContainsAny("hogehoge", "fuga"))
	fmt.Println(strings.ContainsAny("hogehoge", "abc"))
	fmt.Println(strings.ContainsAny("hogehoge", ""))
}

/*
  実行結果
  -------
  true
  true
  false
  true
  false
  false
  -------
*/