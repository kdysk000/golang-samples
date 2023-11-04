package time

import (
	"fmt"
	"time"
)

/*
	経過時間の算出(Since、Seconds)

	書式
	  (time.Since(startTime)).Seconds()
*/
func TimeSample020() {
	fmt.Println("time_sample_020")

	t1 := time.Now()
	t2 := t1.Add(time.Second * -1)  //現在時刻から1秒前

	sub := (time.Since(t2)).Seconds()

	fmt.Println(t1)
	fmt.Println(t2)
	fmt.Println(sub)
}

/*
  実行結果
  -------
  2023-10-21 23:03:33.4817276 +0900 JST m=+0.000101701
  2023-10-21 23:03:32.4817276 +0900 JST m=-0.999898299
  1.0000004
  -------
*/
