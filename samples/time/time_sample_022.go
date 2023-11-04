package time

import (
	"fmt"
	"time"
)

/*
	指定時間後に一度だけ現在時刻を通知する(After)

	指定した時間後に現在時刻を表すtime.Time型の値が送信されるチャネルを生成する

	func After(d Duration) <-chan Time
	  param:
	    d          : 通知の時間間隔
	  return:
	    <-chan Time: Timeオブジェクトが通知されるチャネル
*/
func TimeSample022() {
	fmt.Println("time_sample_022")

	//1秒後に通知
	ch := time.After(1 * time.Second)

	t := <-ch
	fmt.Println(t)
}

/*
  実行結果
  -------
  2023-11-04 11:07:45.7015365 +0900 JST m=+1.000241801
  -------
*/
