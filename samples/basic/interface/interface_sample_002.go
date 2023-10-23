package myinterface

import (
	"fmt"
)

/*
	interface{}型

	・関数の無いインタフェース interface {} は、any 型のように使用することができる
	・関数の引数をinterface{}にするとどんな型の引数でも受け取ることができる
*/
func testFunc1(a interface {}) {
    fmt.Println(a)
}

func InterfaceSample002() {
	fmt.Println("interface_sample_002")

	x := 1
	s := "hello world"

	testFunc1(x)
	testFunc1(s)
}

/*
  実行結果
  -------
  1
  hello world
  -------
*/
