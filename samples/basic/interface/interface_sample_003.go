package myinterface

import (
	"fmt"
)

/*
	interface{} 型の変数を他の型に変換する

	書式
	  .(型名)
*/
func testFunc2(a interface {}) {
	str := a.(string)
	fmt.Printf("%T\n", str)
}

func InterfaceSample003() {
	fmt.Println("interface_sample_003")

	s := "hello world"
	testFunc2(s)
}

/*
  実行結果
  -------
  string
  -------
*/
