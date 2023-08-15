package time

import (
	"fmt"
	"time"
)

/*
	Timeオブジェクトから時刻を取得する(Clock、Hour、Minute、Second)
	func (t Time) Clock() (hour, min, sec int)
		param:
		  なし
		return:
		  hour : 時
		  min  : 分
		  sec  : 秒
	func (t Time) Hour() int
		param:
		  なし
		return:
		  int : 時
	func (t Time) Minute() int
		param:
		  なし
		return:
		  int : 分
	func (t Time) Second() int
		param:
		  なし
		return:
		  int : 秒
*/
func TimeSample004() {
	fmt.Println("time_sample_004")

	t := time.Date(2001, 1, 1, 23, 59, 59, 0, time.UTC)
	hour1, min1, sec1 := t.Clock()

	hour2  := t.Hour()
	min2   := t.Minute()
	sec2   := t.Second()

    fmt.Println(t)
	fmt.Println(hour1)
	fmt.Println(min1)
	fmt.Println(sec1)
	fmt.Println(hour2)
	fmt.Println(min2)
	fmt.Println(sec2)
}

/*
  実行結果
  -------
  2001-01-01 23:59:59 +0000 UTC
  23
  59
  59
  23
  59
  59
  -------
*/
