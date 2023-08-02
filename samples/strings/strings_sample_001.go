package strings

import (
	"fmt"
	"strings"
)

/*
	文字列を比較する(Compare)
	func Compare(a, b string) int
		param:
		  a  : 文字列1
		  b  : 文字列2
		return:
		  int: a == b の場合は 0、 a < b の場合は -1、 a > b の場合は 1 を返す
*/
func StringsSample001() {
	fmt.Println("strings_sample_001")

	fmt.Println(strings.Compare("aaa", "aaa"))
	fmt.Println(strings.Compare("aaa", "bbb"))
	fmt.Println(strings.Compare("bbb", "aaa"))
}
