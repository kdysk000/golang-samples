package strconv

import (
	"fmt"
	"log"
	"strconv"
)

/*
	文字列をbool型に変換する(ParseBool)
	func ParseBool(str string) (bool, error)
		概要:
		  受け付ける値は1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False
		  それ以外はエラー
		param:
		  str   : bool文字列
		return:
		  bool  : boolデータ
		  error : エラー
*/
func StrconvSample002() {
	fmt.Println("strconv_sample_002")

	b, err := strconv.ParseBool("true")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(b)
}

/*
  実行結果
  -------
  true
  -------
*/
