package time

import (
	"fmt"
	"time"
)

/*
	日付の加算減算(AddDate)

	func (t Time) AddDate(years int, months int, days int) Time
	  param:
	    years  : 加算減算する年
	    months : 加算減算する月
	    days   : 加算減算する日
	  return:
	    Time: Timeオブジェクト
*/
func TimeSample011() {
	fmt.Println("time_sample_011")

	t := time.Date(2001, 1, 1, 23, 59, 59, 0, time.UTC)

	afterTenYears  := t.AddDate(10, 0, 0)
	beforeTenYears := t.AddDate(-10, 0, 0)
	afterTenMonths := t.AddDate(0, 10, 0)
	afterTenDays   := t.AddDate(0, 0, 10)

	fmt.Println(t)
	fmt.Println(afterTenYears)
	fmt.Println(beforeTenYears)
	fmt.Println(afterTenMonths)
	fmt.Println(afterTenDays)
}

/*
  実行結果
  -------
  2001-01-01 23:59:59 +0000 UTC
  2011-01-01 23:59:59 +0000 UTC
  1991-01-01 23:59:59 +0000 UTC
  2001-11-01 23:59:59 +0000 UTC
  2001-01-11 23:59:59 +0000 UTC
  -------
*/
