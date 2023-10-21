package time

import (
	"fmt"
	"time"
)

/*
	日時の差分を求める(Since)

	func Since(t Time) Duration
	  概要
	    現在日時と引数で指定された日時の差分を求める
		Sub()と違って開始時点のTImeオブジェクトだけで実行できる
	  param:
	    t       : Timeオブジェクト
	  return:
	    Duration: 差分結果
*/
func TimeSample014() {
	fmt.Println("time_sample_014")

	t1 := time.Now()
	t2 := t1.Add(time.Second * -1)  //現在時刻から1秒前

	sub := time.Since(t2)

	fmt.Println(t1)
	fmt.Println(t2)
	fmt.Println(sub)
}

/*
  実行結果
  -------
  2023-10-21 22:49:43.8646511 +0900 JST m=+0.000097401
  2023-10-21 22:49:42.8646511 +0900 JST m=-0.999902599
  1.0000003s
  -------
*/
