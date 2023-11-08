package sync

import (
	"fmt"
	"sync"
	"time"
)

/*
	RWMutexによる排他制御(RWMutex.RLock、RWMutex.RUnlock)

	sync.MutexはLockのみしか提供しないため、Read/Write双方で使うと効率が悪くなる。
	sync.RWMutexはRLockというもう一つのLockを提供する。

	RLock:
	  Read向けのLock。
	  RLock同士はブロックせずLockのみがブロックされる。RLockの解除はUnlockではなくRUnlockで解除する。
	Lock:
	  RLock, Lock双方をブロックする。

	func (rw *RWMutex) RLock()
	func (rw *RWMutex) RUnlock()
*/

type RWMutexTest struct {
	cnt int
	mux sync.RWMutex
}

func (m *RWMutexTest) CountUp() {
	//ロックして1秒後にカウントアップ
	m.mux.Lock()
	time.Sleep(time.Second)
	m.cnt++
	m.mux.Unlock()
}

func (m *RWMutexTest) Value() int {
	m.mux.RLock()
	defer m.mux.RUnlock()
	time.Sleep(3 * time.Second)
	return m.cnt
}

func SyncSample002() {
	fmt.Println("sync_sample_002")

	m := RWMutexTest{}

	//この2つのゴルーチンはともにRLockしているのでブロックされず、ほぼ同時にログ出力される。
	go func() {
		fmt.Println("[START] FuncA")
		v := m.Value()
		fmt.Println("[END] FuncA", v)
	}()
	go func() {
		fmt.Println("[START] FuncB")
		v := m.Value()
		fmt.Println("[END] FuncB", v)
	}()

	//タイミング調整のためのスリープ
	time.Sleep(time.Second)

	//FuncA,BでRLockされているのでRUnlock()されるまでブロックされる
	go func() {
		fmt.Println("[START] FuncC")
		m.CountUp()
		fmt.Println("[END] FuncC")
	}()

	//タイミング調整のためのスリープ
	time.Sleep(time.Second)

	//FuncCでLockされているのでUnlockされるまでブロックされる
	go func() {
		fmt.Println("[START] FuncD")
		v := m.Value()
		fmt.Println("[END] FuncD", v)
	}()

	time.Sleep(8 * time.Second)
}

/*
  実行結果
  -------
  [START] FuncB
  [START] FuncA
  [START] FuncC
  [START] FuncD
  [END] FuncA 0
  [END] FuncB 0
  [END] FuncC
  [END] FuncD 1
  -------
*/
