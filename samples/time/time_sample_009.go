package time

import (
	"fmt"
	"time"
)

/*
	時刻をUTCに変換する(UTC)

	func (t Time) UTC() Time
	  param:
	    なし
	  return:
	    Time: Timeオブジェクト
*/
func TimeSample009() {
	fmt.Println("time_sample_009")

	t := time.Date(2001, 1, 1, 23, 59, 59, 0, time.Local)
	fmt.Println(t)

	utc := t.UTC()
    fmt.Println(utc)
}

/*
  実行結果
  -------
  2001-01-01 23:59:59 +0900 JST
  2001-01-01 14:59:59 +0000 UTC
  -------
*/
