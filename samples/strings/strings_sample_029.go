package strings

import (
	"fmt"
	"strings"
)

/*
	文字列の先頭と末尾から空白を除去(TrimSpace)

	TrimSpace(s string) string
	  概要
	    sの先頭と末尾から空白を除いた文字列を返す
	  param:
	    s     : 文字列
	    prefix: 除去するプレフィックス
	  return:
	    string: 除去後の文字列
*/
func StringsSample029() {
	fmt.Println("strings_sample_029")

	str := "    Hello world.    "

	str1 := strings.TrimSpace(str)

	fmt.Println(str1)
}

/*
  実行結果
  -------
  Hello world.
  -------
*/
