package time

import (
	"fmt"
	"time"
)

/*
	TimeオブジェクトをUNIXタイムに変換する(time.Unix)
	func (t Time) Unix() int64
		param:
		  なし
		return:
		  int64: UNIXタイム
*/
func TimeSample008() {
	fmt.Println("time_sample_008")

	t := time.Date(2001, 1, 1, 23, 59, 59, 0, time.UTC)
	unixtime := t.Unix()
    fmt.Println(unixtime)
}

/*
  実行結果
  -------
  978393599
  -------
*/
