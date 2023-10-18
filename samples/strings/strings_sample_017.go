package strings

import (
	"fmt"
	"strings"
)

/*
	文字列の結合(Join)

	Join(elems []string, sep string) string
	  概要
	    elemsの要素をsepで結合して1つの文字列を返す
	  param:
	    elms  : 文字列のスライス
	    sep   : セパレータ
	  return:
	    string: 結合した文字列
*/
func StringsSample017() {
	fmt.Println("strings_sample_017")

	elems := []string{"hoge", "fuga", "foo", "bar"}

	str1 := strings.Join(elems, ",")
	str2 := strings.Join(elems, "")

	fmt.Println(str1)
	fmt.Println(str2)
}

/*
  実行結果
  -------
  hoge,fuga,foo,bar
  hogefugafoobar
  -------
*/
