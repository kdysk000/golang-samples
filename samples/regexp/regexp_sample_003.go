package regexp

import (
	"fmt"
	"regexp"
)

/*
	パターンに一致する文字列を取得する(FindString、FindAllString)

	func (re *Regexp) FindString(s string) string
	  概要
	    パターンに一致する最初の文字列を返す
	  param:
	    s     : 文字列
	  return:
	    string: パターンに一致する最初の文字列

	func (re *Regexp) FindAllString(s string, n int) []string
	  概要
	    パターンに一致する文字列の配列を返す
	  param:
	    s       : 文字列
	    n       : 検索結果の最大数(-1を指定すれば全て)
	  return:
	    []string: パターンに一致する文字列の配列
*/

func RegexpSample003() {
	fmt.Println("regexp_sample_003")

	//1個以上の数字文字
	re := regexp.MustCompile(`\d+`)

	str := `123-456-789-0`
	fmt.Println(re.FindString(str))
	fmt.Println(re.FindAllString(str, -1))
}

/*
  実行結果
  -------
  123
  [123 456 789 0]
  -------
*/
