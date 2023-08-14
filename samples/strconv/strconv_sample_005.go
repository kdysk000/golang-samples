package strconv

import (
	"fmt"
	"strconv"
)

/*
	bool型を文字列に変換する(FormatBool)
	func FormatBool(b bool) string
		param:
		  b     : bool型
		return:
		  string: 変換した文字列
*/
func StrconvSample005() {
	fmt.Println("strconv_sample_005")

	fmt.Println(strconv.FormatBool(true))
	fmt.Println(strconv.FormatBool(false))
}

/*
  実行結果
  -------
  true
  false
  -------
*/
