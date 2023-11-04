package time

import (
	"fmt"
	"time"
)

/*
	タイムゾーンの変更(In)

	func (t Time) In(loc *Location) Time
	  param:
	    loc : タイムゾーン
	  return:
	    Time: Timeオブジェクト
*/
func TimeSample011() {
	fmt.Println("time_sample_011")

	t1 := time.Unix(0,0)
	t2 := t1.In(time.UTC)
    fmt.Println(t1)
	fmt.Println(t2)
}

/*
  実行結果
  -------
  1970-01-01 09:00:00 +0900 JST
  1970-01-01 00:00:00 +0000 UTC
  -------
*/
