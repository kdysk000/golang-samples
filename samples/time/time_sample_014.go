package time

import (
	"fmt"
	"time"
)

/*
	日時の比較(Equal、After、Before)
	func (t Time) Equal(u Time) bool
		param:
		  u   : 比較するTimeオブジェクト
		return:
		  bool: 比較結果
	func (t Time) After(u Time) bool
		param:
		  u   : 比較するTimeオブジェクト
		return:
		  bool: 比較結果(uよりも後の日時ならtrue)
	func (t Time) Before(u Time) bool
		param:
		  u   : 比較するTimeオブジェクト
		return:
		  bool: 比較結果(uよりも前の日時ならtrue)
*/
func TimeSample014() {
	fmt.Println("time_sample_014")

	t1 := time.Date(2001, 1, 1, 23, 59, 59, 0, time.UTC)
	t2 := time.Date(2001, 1, 1, 23, 59, 58, 0, time.UTC)

	fmt.Println(t1.Equal(t2))
	fmt.Println(t1.After(t2))
	fmt.Println(t1.Before(t2))
}

/*
  実行結果
  -------
  false
  true
  false
  -------
*/
