package time

import (
	"fmt"
	"time"
)

/*
	指定の時間停止させる(Sleep)

	func Sleep(d Duration)
	  param:
	    d: 停止させる時間
	  return:
	    なし
*/
func TimeSample018() {
	fmt.Println("time_sample_018")

	fmt.Println(time.Now())
	time.Sleep(time.Second * 1)
	fmt.Println(time.Now())
}

/*
  実行結果
  -------
  2023-11-15 22:41:40.9653826 +0900 JST m=+0.000102501
  2023-11-15 22:41:41.9657496 +0900 JST m=+1.000469401
  -------
*/