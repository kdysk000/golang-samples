package strconv

import (
	"fmt"
	"strconv"
)

/*
	数値を文字列に変換する(Itoa)

	func Itoa(i int) string
	  param:
	    i     : int型
	  return:
	    string: 変換した数値文字列
*/
func StrconvSample004() {
	fmt.Println("strconv_sample_004")

	fmt.Println(strconv.Itoa(255))
	fmt.Println(strconv.Itoa(-1))
}

/*
  実行結果
  -------
  255
  -1
  -------
*/
