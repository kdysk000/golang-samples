package sync

import (
	"fmt"
	"sync"
	"time"
)

/*
	Mutexによる排他制御(Mutex.Lock、Mutex.Unlock)

	複数のゴルーチンが動作している場合、同じリソースに同時にアクセスし処理が衝突する可能性がある。
	mutexを使用すると1つのゴルーチンだけが処理コードにアクセスできるようにロックして処理の衝突を防ぐことができる。

	func (m *Mutex) Lock()
	  概要：
	    ロックする。
	    すでにロックされている場合、呼び出し元のゴルーチンは、ミューテックスが使用可能になるまでブロックされる。

	func (m *Mutex) TryLock() bool
	  概要
	    Lock()との違いはすでにロックされている場合、ブロックされない。
	    ロックできた場合はtrueが返り、ロックできなかった場合はfalseが返る。

	func (m *Mutex) Unlock()
	  概要：
	    ロック解除する。
	    ロックされていない状態で実行するとエラーになる
*/

type MutexTest struct {
	cnt int
	mux sync.Mutex
}

func (m *MutexTest) CountUp() {
	//ロックして1秒後にカウントアップ
	m.mux.Lock()
	time.Sleep(time.Second)
	m.cnt++
	m.mux.Unlock()
}

func (m *MutexTest) Value() int {
	m.mux.Lock()
	defer m.mux.Unlock()
	return m.cnt
}

func SyncSample001() {
	fmt.Println("sync_sample_001")

	m := MutexTest{}

	//以下の2つのゴルーチンどちらが先に実行されるかはわからないが、
	//あとのゴルーチンは先のゴルーチンがロック解除するまで待ってからカウントアップされる
	go func() {
		m.CountUp()
	}()

	go func() {
		m.CountUp()
	}()

	time.Sleep(3 * time.Second)
	fmt.Println(m.Value())
}

/*
  実行結果
  -------
  2
  -------
*/
