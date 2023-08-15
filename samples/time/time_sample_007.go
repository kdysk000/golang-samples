package time

import (
	"fmt"
	"time"
)

/*
	UNIXタイムをTimeオブジェクトに変換する(time.Unix)
	func Unix(sec int64, nsec int64) Time
		param:
		  sec : 秒
          nsec: ナノ秒
		return:
		  Time: Timeオブジェクト
*/
func TimeSample007() {
	fmt.Println("time_sample_007")

	t := time.Unix(0,0)
    fmt.Println(t)
}

/*
  実行結果
  -------
  1970-01-01 09:00:00 +0900 JST
  -------
*/
