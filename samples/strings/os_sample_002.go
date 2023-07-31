package strings

import (
	"fmt"
	"strings"
)

/*
	指定の文字列を含んでいるか(Contains)
	func Contains(s, substr string) bool
		param:
		  s     : 検索対象の文字列
		  substr: 検索する文字列
		return:
		  bool: 
		注：
		  substrが空文字列ならtrueが返る
*/
func StringsSample002() {
	fmt.Println("strings_sample_002")

	fmt.Println(strings.Contains("hogehoge", "h"))
	fmt.Println(strings.Contains("hogehoge", "hoge"))
	fmt.Println(strings.Contains("hogehoge", "fuga"))
	fmt.Println(strings.Contains("hogehoge", ""))
	fmt.Println(strings.Contains("", ""))
}
