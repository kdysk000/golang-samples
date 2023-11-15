package regexp

import (
	"fmt"
	"regexp"
)

/*
	パターンに一致する文字があるかどうかを調べる(MatchString)

	func (re *Regexp) MatchString(s string) bool
	  param:
	    s   : 文字列
	  return:
	    bool: パターンに一致すればtrue、しなければfalse
*/

func RegexpSample002() {
	fmt.Println("regexp_sample_002")

	//1個以上の数字文字
	re := regexp.MustCompile(`\d+`)

	str1 := `123-456-789-0`
	fmt.Println(re.MatchString(str1))

	str2 := `abc-def-ghu-j`
	fmt.Println(re.MatchString(str2))
}

/*
  実行結果
  -------
  true
  false
  -------
*/
