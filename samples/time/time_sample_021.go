package time

import (
	"fmt"
	"time"
)

/*
	指定時間ごとに現在時刻を通知する(Tick)

	指定した時間間隔ごとに現在時刻を表すtime.Time型の値が送信されるチャネルを生成する

	func Tick(d Duration) <-chan Time
	  param:
	    d          : 通知の時間間隔
	  return:
	    <-chan Time: Timeオブジェクトが通知されるチャネル
*/
func TimeSample021() {
	fmt.Println("time_sample_021")

	//1秒ごとに通知
	ch := time.Tick(1 * time.Second)

	for i := 0; i < 3; i++ {
		t := <-ch
		fmt.Println(t)
	}
}

/*
  実行結果
  -------
  2023-11-04 11:00:47.2809684 +0900 JST m=+1.000430301
  2023-11-04 11:00:48.2814435 +0900 JST m=+2.000905501
  2023-11-04 11:00:49.2809562 +0900 JST m=+3.000418201
  -------
*/
