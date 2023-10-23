package myinterface

import (
	"fmt"
)

/*
	インターフェース(Interface)の定義

	インターフェースとは
	  メソッドの型だけを列挙した型のこと。
	  あるインターフェース Foo で宣言されているメソッドがすべて実装されている構造体であれば何でも型 Foo として扱うことができる。
	  インタフェースはポリモーフィズムを実装するための機能。

	書式
	  type 型名 interface {
	  	メソッド名1(引数の型, ...) (返り値の型, ...)
	  	.....
	  	メソッド名n(引数の型, ...) (返り値の型, ...)
	  }
*/
type human interface {
	greet()
}

type japanese struct {}
type american struct {}

func (j japanese) greet() {
	fmt.Println("こんにちは")
}

func (a american) greet() {
	fmt.Println("Hello")
}

func InterfaceSample001() {
	fmt.Println("interface_sample_001")

	var h human
	h = japanese{}  //interfaceの型に代入
	h.greet()

	h = american{}  //interfaceの型に代入
	h.greet()
}

/*
  実行結果
  -------
  こんにちは
  Hello
  -------
*/
