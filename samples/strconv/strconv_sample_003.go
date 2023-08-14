package strconv

import (
	"fmt"
	"strconv"
)

/*
	文字列を数値に変換する(ParseInt、ParseUint、ParseFloat)
	func ParseInt(s string, base int, bitSize int) (i int64, err error)
		param:
		  s      : 数値文字列
		  base   : 変換に用いる基数(2〜36)。0の場合s文字列の書式から判断する(0x接頭詞がついていたら16進数など)
		  bitsize: 0,8,16,32,64(それぞれint, int8, int16, int32, uint64に該当)
		return:
		  i      : 変換した数値
		  err    : エラー
	func ParseUint(s string, base int, bitSize int) (n uint64, err error)
		param:
		  s      : 数値文字列
		  base   : 変換に用いる基数(2〜36)。0の場合s文字列の書式から判断する(0x接頭詞がついていたら16進数など)
		  bitsize: 0,8,16,32,64(それぞれuint, uint8, uint16, uint32, uint64に該当)
		return:
		  n      : 変換した数値
		  err    : エラー
	func ParseFloat(s string, bitSize int) (f float64, err error)
		param:
		  s      : 数値文字列
		  bitsize: 32,64(float32, float64に該当)
		return:
		  f      : 変換した数値
		  err    : エラー
*/
func StrconvSample003() {
	fmt.Println("strconv_sample_003")

	i1, _ := strconv.ParseInt("255", 10, 32)
	i2, _ := strconv.ParseInt("ff", 16, 16)
	i3, _ := strconv.ParseInt("0xff", 0, 16)

	u1, _ := strconv.ParseUint("255", 10, 32)
	u2, _ := strconv.ParseUint("ff", 16, 16)
	u3, _ := strconv.ParseUint("0xff", 0, 16)

	f1, _ := strconv.ParseFloat("3.1415", 32)
	f2, _ := strconv.ParseFloat("3.1415", 64)

	fmt.Println(i1)
	fmt.Println(i2)
	fmt.Println(i3)
	fmt.Println(u1)
	fmt.Println(u2)
	fmt.Println(u3)
	fmt.Println(f1)
	fmt.Println(f2)
}

/*
  実行結果
  -------
  255
  255
  255
  255
  255
  255
  3.1414999961853027
  3.1415
  -------
*/
