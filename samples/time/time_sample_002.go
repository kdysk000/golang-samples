package time

import (
	"fmt"
	"time"
)

/*
	時刻オブジェクトを作成する(time.Date)
	func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time
		param:
		  year : 年
		  month: 月
		  day  : 日
		  hour : 時
		  min  : 分
		  sec  : 秒
		  nsec : ナノ秒
		  loc  : Location(time.UTC、time.Local)
		return:
		  Time: Timeオブジェクト
*/
func TimeSample002() {
	fmt.Println("time_sample_002")

	t1 := time.Date(2001, 1, 1, 23, 59, 59, 0, time.UTC)
	t2 := time.Date(2001, 1, 1, 23, 59, 59, 0, time.Local)

	fmt.Println(t1)
	fmt.Println(t2)
}

/*
  実行結果
  -------
  2001-01-01 23:59:59 +0000 UTC
  2001-01-01 23:59:59 +0900 JST
  -------
*/
