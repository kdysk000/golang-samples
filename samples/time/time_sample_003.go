package time

import (
	"fmt"
	"time"
)

/*
	Timeオブジェクトから日時を取得する(Date、Year、Month、Day、Yearday、Weekday)
	func (t Time) Date() (year int, month Month, day int)
		param:
		  なし
		return:
		  year : 年
		  month: 月
		  day  : 日
	func (t Time) Year() int
		param:
		  なし
		return:
		  int : 年
	func (t Time) Month() Month
		param:
		  なし
		return:
		  Month : 月
	func (t Time) Day() int
		param:
		  なし
		return:
		  int : 日
	func (t Time) YearDay() int
		param:
		  なし
		return:
		  int : 年内通算日数
	func (t Time) Weekday() Weekday
		param:
		  なし
		return:
		  Weekday : 曜日(time.Sunday==0、time.Monday==1、・・・)
*/
func TimeSample003() {
	fmt.Println("time_sample_003")

	t := time.Date(2001, 12, 31, 23, 59, 59, 0, time.UTC)
	year1, month1, day1 := t.Date()

	year2  := t.Year()
	month2 := t.Month()
	day2   := t.Day()
	yd     := t.YearDay()
	wd     := t.Weekday()


    fmt.Println(t)
	fmt.Println(year1)
	fmt.Println(month1)
	fmt.Println(day1)
	fmt.Println(year2)
	fmt.Println(month2)
	fmt.Println(day2)
	fmt.Println(yd)
	fmt.Println(wd)
}

/*
  実行結果
  -------
  2001-12-31 23:59:59 +0000 UTC
  2001
  December
  31
  2001
  December
  31
  365
  Monday
  -------
*/
