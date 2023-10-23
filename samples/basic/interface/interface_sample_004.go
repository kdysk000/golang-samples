package myinterface

import (
	"fmt"
)

/*
	interface{} 型の変数を元の型に変換できるか調べる

	書式
	  q, ok := a.(型名)

	第二戻り値が型変換可能かどうかを返す
*/
type testInterface interface {
	testPrint(i int)
}

type testStruct1 struct {}
type testStruct2 struct {}

// インターフェースを実装している
func (t testStruct1) testPrint(i int) {
	fmt.Println("i:", i)
}

// インターフェースを実装していない
func (t testStruct2) testPrint2(i int) {
	fmt.Println("i:", i)
}

func testFunc3(a interface {}) {
	//aがtestStruct1なら変換可能(okがtrueになる)
	_, ok := a.(testStruct1)
	fmt.Println("ok:", ok)
}

func InterfaceSample004() {
	fmt.Println("interface_sample_004")

	testFunc3(testStruct1{})
	testFunc3(testStruct2{})
}

/*
  実行結果
  -------
  ok: true
  ok: false
  -------
*/
