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
func TimeSample015() {
	fmt.Println("time_sample_015")

	fmt.Println(time.Now())
	time.Sleep(time.Second * 1)
	fmt.Println(time.Now())
}
