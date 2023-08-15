package time

import (
	"fmt"
	"time"
)

/*
	日時の差分を求める(Sub)
	func (t Time) Sub(u Time) Duration
		param:
		  u       : 差分を求めたいTimeオブジェクト
		return:
		  Duration: 差分結果
*/
func TimeSample013() {
	fmt.Println("time_sample_013")

	t1 := time.Date(2001, 1, 1, 23, 59, 59, 0, time.UTC)
	t2 := time.Date(2001, 1, 15, 23, 59, 59, 0, time.UTC)

	sub := t2.Sub(t1)

	fmt.Println(t1)
	fmt.Println(t2)
	fmt.Println(sub)
}

/*
  実行結果
  -------
  2001-01-01 23:59:59 +0000 UTC
  2001-01-15 23:59:59 +0000 UTC
  336h0m0s
  -------
*/
