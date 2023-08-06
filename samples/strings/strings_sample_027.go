package strings

import (
	"fmt"
	"strings"
)

/*
	文字列から指定のプレフィックスを除去(TrimPrefix)
	TrimPrefix(s, prefix string) string
		概要
		  sから接頭辞prefixを除いた文字列を返す
		  sがprefixから始まらない場合はsをそのまま返す
		param:
		  s     : 文字列
		  prefix: 除去するプレフィックス
		return:
		  string: 除去後の文字列
*/
func StringsSample027() {
	fmt.Println("strings_sample_027")

	str := "@! Hello world. @@@!! "

	str1 := strings.TrimPrefix(str, "@! ")
	str2 := strings.TrimPrefix(str, "!@ ")

	fmt.Println(str1)
	fmt.Println(str2)
}

/*
  実行結果
  -------
  Hello world. @@@!! 
  @! Hello world. @@@!! 
  -------
*/
