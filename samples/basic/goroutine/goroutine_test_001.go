package myinterface

import (
	"fmt"
	"time"
)

/*
	ゴルーチン(Goroutine)の定義

	ゴルーチンとは
	  ・Go言語で並行処理を実現するもの
	  ・スレッドよりも簡単にかつ高速に並行処理が実行できる

	書式
	  go funcA()
*/
func funcA() {
	for i := 0; i < 5; i++ {
		fmt.Println("Sub:", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func GoroutineSample001() {
	fmt.Println("goroutine_sample_001")

	//ゴルーチン実行
	go funcA()

	for i := 0; i < 5; i++ {
		fmt.Println("Main:", i)
		time.Sleep(100 * time.Millisecond)
	}
}

/*
  実行結果(順序は実行ごとに変わる可能性あり)
  -------
  Main: 0
  Sub: 0
  Main: 1
  Sub: 1
  Sub: 2
  Main: 2
  Main: 3
  Sub: 3
  Sub: 4
  Main: 4
  -------
*/
