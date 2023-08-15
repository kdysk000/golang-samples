package time

import (
	"fmt"
	"time"
)

/*
	月初、月末を求める(Date、AddDate)
*/
func TimeSample012() {
	fmt.Println("time_sample_012")

	t1 := time.Date(2001, 1, 15, 23, 59, 59, 0, time.UTC)

	// 月初
	t2 := time.Date(t1.Year(), t1.Month(), 1, 0, 0, 0, 0, time.UTC)

	// 月末(月初 + 1month - 1day)
	t3 := t2.AddDate(0, 1, -1)

	fmt.Println(t1)
	fmt.Println(t2)
	fmt.Println(t3)
}

/*
  実行結果
  -------
  2001-01-15 23:59:59 +0000 UTC
  2001-01-01 00:00:00 +0000 UTC
  2001-01-31 00:00:00 +0000 UTC
  -------
*/
