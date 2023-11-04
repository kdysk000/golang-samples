package time

import (
	"fmt"
	"time"
)

/*
	経過時間の算出(Sub、Seconds)

	書式
	  (endTime.Sub(startTime)).Seconds()

	Durationはfloat64型に変換するメソッドを持っているのでこれを使うことで経過時間を数値で取得することができる
	  ・Seconds：秒
	  ・Minutes：分
	  ・Milliseconds：ミリ秒
	  ・Nanoseconds ：ナノ秒

	func (d Duration) Seconds() float64
	  param:
	    なし
	  return:
	    float64: 経過時間(秒)
*/
func TimeSample019() {
	fmt.Println("time_sample_019")

	t1 := time.Date(2001, 1, 1, 23, 59, 58, 0, time.UTC)
	t2 := time.Date(2001, 1, 1, 23, 59, 59, 0, time.UTC)

	sub := (t2.Sub(t1)).Seconds()

	fmt.Println(t1)
	fmt.Println(t2)
	fmt.Println(sub)
}

/*
  実行結果
  -------
  2001-01-01 23:59:58 +0000 UTC
  2001-01-01 23:59:59 +0000 UTC
  1
  -------
*/
