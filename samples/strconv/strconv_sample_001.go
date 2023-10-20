package strconv

import (
	"fmt"
	"log"
	"strconv"
)

/*
	文字列を数値(int)に変換する(Atoi)

	func Atoi(s string) (int, error)
	  param:
	    s    : 数値文字列
	  return:
	    int  : 変換したint型整数
	    error: エラー
*/
func StrconvSample001() {
	fmt.Println("strconv_sample_001")

	num, err := strconv.Atoi("100")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(num)
}

/*
  実行結果
  -------
  100
  -------
*/
