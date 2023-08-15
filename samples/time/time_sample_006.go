package time

import (
	"fmt"
	"time"
)

/*
	時刻オブジェクトを文字列に変換する(String)
	func (t Time) String() string
		param:
		  なし
		return:
		  string: 時刻文字列
*/
func TimeSample006() {
	fmt.Println("time_sample_006")

	t := time.Now()
    fmt.Println(t.String())
}

/*
  実行結果
  -------
  2023-08-14 13:37:52.5743885 +0900 JST m=+0.000088401
  2023
  August
  14
  13
  37
  52
  Monday
  -------
*/
