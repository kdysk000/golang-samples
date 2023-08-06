package strings

import (
	"fmt"
	"strings"
)

/*
	文字列から指定のサフィックスを除去(TrimSuffix)
	TrimSuffix(s, suffix string) string
		概要
		  sから接頭辞suffixを除いた文字列を返す
		  sがsuffixから始まらない場合はsをそのまま返す
		param:
		  s     : 文字列
		  suffix: 除去するサフィックス
		return:
		  string: 除去後の文字列
*/
func StringsSample028() {
	fmt.Println("strings_sample_028")

	str := "@! Hello world. @@@!! "

	str1 := strings.TrimSuffix(str, " @@@!! ")
	str2 := strings.TrimSuffix(str, " @@@!!")

	fmt.Println(str1)
	fmt.Println(str2)
}

/*
  実行結果
  -------
  @! Hello world.
  @! Hello world. @@@!! 
  -------
*/
