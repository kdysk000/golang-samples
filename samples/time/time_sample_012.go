package time

import (
	"fmt"
	"time"
)

/*
	時刻の加算減算(Add)

	func (t Time) Add(d Duration) Time
	  param:
	    d : 加算減算する時刻
	  return:
	    Time: Timeオブジェクト
*/
func TimeSample012() {
	fmt.Println("time_sample_012")

	t := time.Date(2001, 1, 1, 23, 59, 59, 0, time.UTC)

	afterTenSeconds  := t.Add(time.Second * 10)
	beforeTenSeconds := t.Add(time.Second * -10)
	afterTenMinutes  := t.Add(time.Minute * 10)
	afterTenHours    := t.Add(time.Hour * 10)
	afterTenDays     := t.Add(time.Hour * 24 * 10)

	fmt.Println(t)
	fmt.Println(afterTenSeconds)
	fmt.Println(beforeTenSeconds)
	fmt.Println(afterTenMinutes)
	fmt.Println(afterTenHours)
	fmt.Println(afterTenDays)
}

/*
  実行結果
  -------
  2001-01-01 23:59:59 +0000 UTC
  2001-01-02 00:00:09 +0000 UTC
  2001-01-01 23:59:49 +0000 UTC
  2001-01-02 00:09:59 +0000 UTC
  2001-01-02 09:59:59 +0000 UTC
  2001-01-11 23:59:59 +0000 UTC
  -------
*/
