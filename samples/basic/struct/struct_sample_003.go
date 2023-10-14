package mystruct

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
type testStruct3 struct {
	x int
	y int
}

func StructSample003() {
	fmt.Println("struct_sample_003")

	st := testStruct3{5,10}
	fmt.Println(st.getNum())
}

func (s testStruct3) getNum() (int, int) {
	return s.x, s.y
}

/*
  実行結果
  -------
  5 10
  -------
*/
