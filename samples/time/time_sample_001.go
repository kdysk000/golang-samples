package time

import (
	"fmt"
	"time"
)

/*
	現在時刻を取得する(time.Now)
	func Now() Time
		param:
		  なし
		return:
		  Time: 現在時刻データを保持するTimeオブジェクト
*/
func TimeSample001() {
	fmt.Println("time_sample_001")

	t := time.Now()
    fmt.Println(t)
}

/*
  実行結果
  -------
  2023-08-14 13:37:52.5743885 +0900 JST m=+0.000088401
  -------
*/
