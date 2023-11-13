package sync

import (
	"fmt"
	"sync"
	"time"
)

/*
	WaitGroupによるゴルーチンの完了待ち(WaitGroup.Add、WaitGroup.Done、WaitGroup.Wait)

	全てのゴルーチンが完了するまで待ちたい場合などに使用

	func (wg *WaitGroup) Add(delta int)
	  待ち受けるゴルーチンの数を設定

	func (wg *WaitGroup) Done()
	  ゴルーチンの完了を設定

	func (wg *WaitGroup) Wait()
	  Addした数分Doneされるまで待つ
*/

func SyncSample003() {
	fmt.Println("sync_sample_003")

	var wg sync.WaitGroup

	//待ち受けするゴルーチンの数を設定
	wg.Add(2)

	go func() {
		fmt.Println("[START] FuncA")
		time.Sleep(time.Second)
		fmt.Println("[END] FuncA")
		wg.Done()  //完了
	}()
	go func() {
		fmt.Println("[START] FuncB")
		time.Sleep(time.Second)
		fmt.Println("[END] FuncB")
		wg.Done()  //完了
	}()

	//ゴルーチンの完了を待つ
	wg.Wait()

	fmt.Println("All Goroutine finish.")
}

/*
  実行結果
  -------
  [START] FuncB
  [START] FuncA
  [END] FuncA
  [END] FuncB
  All Goroutine finish.
  -------
*/
