package strings

import (
	"fmt"
	"strings"
)

/*
	指定の文字列が含まれている個数を取得(Count)

	Count(s, substr string) int
	  概要
	    s に含まれている substr の数をカウントして返す
	    substr が空文字列の場合、 s に含まれている Unicode コードポイントの数 + 1 を返す
	  param:
	    s     : 検索対象の文字列
	    string: 検索する文字列
	  return:
	    int   : 見つかった数
*/
func StringsSample005() {
	fmt.Println("strings_sample_005")

	fmt.Println(strings.Count("hogehoge", "h"))
	fmt.Println(strings.Count("hogehoge", "hoge"))
	fmt.Println(strings.Count("hogehoge", "HOGE"))
	fmt.Println(strings.Count("hogehoge", "fuga"))
	fmt.Println(strings.Count("hogehoge", ""))
}

/*
  実行結果
  -------
  2
  2
  0
  0
  9
  -------
*/
