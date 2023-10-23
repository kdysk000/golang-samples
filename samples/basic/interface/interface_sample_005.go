package myinterface

import (
	"fmt"
)

/*
	interface{} 型の変数の元の型を判断する

	書式
	  switch 変数.(type) { ... }
*/
func testFunc4(a interface {}) {
	switch a.(type) {
	case int:
		fmt.Println("a is integer")
	case string:
		fmt.Println("a is string")
	default:
		fmt.Println("a is other type")
	}
}

func InterfaceSample005() {
	fmt.Println("interface_sample_005")

	i := 1
	s := "hello world"
	b := true

	testFunc4(i)
	testFunc4(s)
	testFunc4(b)
}

/*
  実行結果
  -------
  a is integer
  a is string
  a is other type
  -------
*/
