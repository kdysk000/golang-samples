package strconv

import (
	"fmt"
	"strconv"
)

/*
	数値を文字列に変換する(FormatInt、FormatUint、FormatFloat)
	func FormatInt(i int64, base int) string
		param:
		  i     : int64型
		  base  : 基数(2〜36)
		return:
		  string: 数値文字列
	func FormatUint(i uint64, base int) string
		param:
		  i     : uint64型
		  base  : 基数(2〜36)
		return:
		  string: 数値文字列
	func FormatFloat(f float64, fmt byte, prec, bitSize int) string
		param:
		  f      : float64型
		  fmt    : 出力書式
					bは二進対数表記
					eとEは十進対数表記
					fは非対数表記
					g,Gは大きい数は対数、そうでなければ非対数的に表記
		  prec   : 桁数
					eEFのとき小数点以下桁数
					gGのとき全体の桁数
		  bitsize: 32,64(float32, float64に該当)
		return:
		  string: 数値文字列
*/
func StrconvSample006() {
	fmt.Println("strconv_sample_006")

	fmt.Println(strconv.FormatInt(255, 10))
	fmt.Println(strconv.FormatInt(255, 16))
	fmt.Println(strconv.FormatInt(-1, 10))

	fmt.Println(strconv.FormatUint(255, 10))
	fmt.Println(strconv.FormatUint(255, 16))

	fmt.Println(strconv.FormatFloat(1234.56789, 'b', 5, 64))
	fmt.Println(strconv.FormatFloat(1234.56789, 'e', 5, 64))
	fmt.Println(strconv.FormatFloat(1234.56789, 'E', 5, 64))
	fmt.Println(strconv.FormatFloat(1234.56789, 'f', 5, 64))
	fmt.Println(strconv.FormatFloat(1234.56789, 'g', 5, 64))
	fmt.Println(strconv.FormatFloat(1234.56789, 'G', 5, 64))
	fmt.Println(strconv.FormatFloat(1234.56789, 'G', 4, 64))
}

/*
  実行結果
  -------
  255
  ff
  -1
  255
  ff
  5429687001335527p-42
  1.23457e+03
  1.23457E+03
  1234.56789
  1234.6
  1234.6
  1235
  -------
*/
