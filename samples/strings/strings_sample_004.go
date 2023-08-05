package strings

import (
	"fmt"
	"strings"
)

/*
	指定のUnicodeコードポイントが含まれているか(ContainsRune)
	ContainsRune(s string, r rune) bool
		param:
		  s  : 検索対象の文字列
		  r  : 検索するUnicodeコードポイント
		return:
		  bool:
*/
func StringsSample004() {
	fmt.Println("strings_sample_004")

	fmt.Println(strings.ContainsRune("hogehoge", 'h'))
	fmt.Println(strings.ContainsRune("hogehoge", 'H'))
}

/*
  実行結果
  -------
  true
  false
  -------
*/