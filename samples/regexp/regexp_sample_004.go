package regexp

import (
	"fmt"
	"regexp"
)

/*
	パターンに一致する文字列の位置を取得する(FindStringIndex、FindAllStringIndex)

	func (re *Regexp) FindStringIndex(s string) (loc []int)
	  概要
	    パターンに一致する最初の文字列の始まりと終わりの位置を返す
	  param:
	    s  : 文字列
	  return:
	    loc: 位置

	func (re *Regexp) FindAllStringIndex(s string, n int) [][]int
	  概要
	    パターンに一致する最初の文字列の始まりと終わりの位置を返す
	  param:
	    s      : 文字列
		n      : 検索結果の最大数(-1を指定すれば全て)
	  return:
	    [][]int: 位置の配列
*/

func RegexpSample004() {
	fmt.Println("regexp_sample_004")

	//1個以上の数字文字
	re := regexp.MustCompile(`\d+`)

	str := `123-456-789-0`
	fmt.Println(re.FindStringIndex(str))
	fmt.Println(re.FindAllStringIndex(str, -1))
}

/*
  実行結果
  -------
  [0 3]
  [[0 3] [4 7] [8 11] [12 13]]
  -------
*/
