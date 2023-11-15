package regexp

import (
	"fmt"
	"regexp"
)

/*
	正規表現のパターンを定義する(Compile、MustCompile)

	func Compile(expr string) (*Regexp, error)
	  param:
	    expr   : パターン文字列
	  return:
	    *Regexp:
	    error  : エラー

	func MustCompile(str string) *Regexp
	  param:
	    expr   : パターン文字列
	  return:
	    *Regexp:

	  MustCompile()はコンパイルエラー時にエラーを返すのではなくパニックを発生させる
*/

func RegexpSample001() {
	fmt.Println("regexp_sample_001")

	re1, _ := regexp.Compile(`ab+c`)
	fmt.Println(re1)

	re2 := regexp.MustCompile(`ab+c`)
	fmt.Println(re2)
}

/*
  実行結果
  -------
  ab+c
  ab+c
  -------
*/
