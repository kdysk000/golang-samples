package basic

import (
	"fmt"
)

/*
	構造体のメソッド

	書式
	    func レシーバ メソッド名(引数リスト) 戻り値リスト {
			// do somthing
		}

	下記 (s testStruct2) の s はレシーバと呼ばれる。
	関数の定義にレシーバを指定することによって、レシーバで指定した型(構造体)と関数を関連付ける。
*/
type testStruct2 struct {
    x int
    y int
}

func BasicSample027() {
	fmt.Println("basic_sample_027")

	st := testStruct2{5,10}
	fmt.Println(st.getNum())
}

func (s testStruct2) getNum() (int, int) {
    return s.x, s.y
}

/*
  実行結果
  -------
  5 10
  -------
*/
