package time

import (
	"fmt"
	"time"
)

/*
	Timeオブジェクトを文字列に変換する(String)

	func (t Time) String() string
	  param:
	    なし
	  return:
	    string: 時刻文字列
*/
func TimeSample006() {
	fmt.Println("time_sample_006")

	t := time.Date(2001, 1, 1, 23, 59, 59, 0, time.UTC)
    fmt.Println(t.String())
}

/*
  実行結果
  -------
  2001-01-01 23:59:59 +0000 UTC
  -------
*/
