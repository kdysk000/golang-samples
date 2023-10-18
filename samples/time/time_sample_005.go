package time

import (
	"fmt"
	"time"
)

/*
	指定フォーマットで時刻文字列を取得する(Format)

	func (t Time) Format(layout string) string
	  param:
	    layout: "Mon Jan 2 15:04:05 -0700 MST 2006" を並べ替えた文字列
	  return:
	    string: 現在時刻データを保持するTimeオブジェクト
*/
func TimeSample005() {
	fmt.Println("time_sample_005")

    t := time.Now()
    const layout = "2006-01-02 15:04:05 MST"
    fmt.Println(t.Format(layout))
}

/*
  実行結果
  -------
  2023-08-14 14:19:49 JST
  -------
*/
