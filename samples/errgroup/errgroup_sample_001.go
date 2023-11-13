package errgroup

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/sync/errgroup"
)

/*
	errgroupによるゴルーチンの完了待ち(Go、Wait)

	ゴルーチンのエラーハンドリングを行いたい時に使用する。
	いずれかの goroutine でエラーがあったら最初のひとつを知ることができる。(つまり全てのゴルーチンが成功したかどうかを知ることができる)
	ctx を組み合わせても使えるようになっているので、goroutine のどれかがエラーになったら処理を切り上げる、という使い方も可能。

	func (g *Group) Go(f func() error)
	  ゴルーチンの実行

	func (g *Group) Wait() error
	  いずれかのゴルーチンでエラーが発生した場合、最初のエラーを返す
*/

func ErrgroupSample001() {
	fmt.Println("errgroup_sample_001")

	funcs := []string{
		"FuncA",
		"FuncB",
		"FuncC",
	}

	testFunc := func(name string) func() error {
		return func() error {
			fmt.Println("[START]", name)
			time.Sleep(time.Second)
			fmt.Println("[END]", name)
			return nil
		}
	}

	eg := errgroup.Group{}

	for _, f := range funcs {
		f := f
		//ゴルーチン起動
		eg.Go(testFunc(f))
	}

	//全てのゴルーチンが完了するのを待つ
	if err := eg.Wait(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("All Goroutine finish.")
}

/*
  実行結果
  -------
  [START] FuncC
  [START] FuncA
  [START] FuncB
  [END] FuncB
  [END] FuncA
  [END] FuncC
  All Goroutine finish.
  -------
*/
