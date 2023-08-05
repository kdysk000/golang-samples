package strings

import (
	"fmt"
	"strings"
)

/*
	文字列を大文字小文字区別せずに比較(EqualFold)
	EqualFold(s, t string) bool
		param:
		  s   : 文字列1
		  t   : 文字列2
		return:
		  bool:
*/
func StringsSample006() {
	fmt.Println("strings_sample_006")

	fmt.Println(strings.EqualFold("hogehoge", "h"))
	fmt.Println(strings.EqualFold("hogehoge", "hogehoge"))
	fmt.Println(strings.EqualFold("hogehoge", "HOGEHOGE"))
	fmt.Println(strings.EqualFold("hogehoge", "HoGehOgE"))
}

/*
  実行結果
  -------
  false
  true
  true
  true
  -------
*/
