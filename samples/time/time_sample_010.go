package time

import (
	"fmt"
	"time"
)

/*
	時刻をローカルタイムに変換する(Local)

	func (t Time) Local() Time
	  param:
	    なし
	  return:
	    Time: Timeオブジェクト
*/
func TimeSample010() {
	fmt.Println("time_sample_010")

	t := time.Date(2001, 1, 1, 23, 59, 59, 0, time.UTC)
	fmt.Println(t)

	local := t.Local()
    fmt.Println(local)
}

/*
  実行結果
  -------
  2001-01-01 23:59:59 +0000 UTC
  2001-01-02 08:59:59 +0900 JST
  -------
*/
