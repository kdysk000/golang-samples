package mystruct

import (
	"fmt"
)

/*
	構造体の埋め込み

	Go言語は他のオブジェクト指向言語のようなクラスがない
	継承に似た機能として構造体の埋め込みがある

*/
type testStruct5 struct {
	x int
	y int
}

type testStruct6 struct {
	testStruct5  //testStruct5を埋め込み
}

func StructSample005() {
	fmt.Println("struct_sample_005")

	st := testStruct6{}
	st.setNum(5, 10)
	fmt.Println(st.x, st.y)
}

func (s *testStruct6) setNum(x int, y int) {
	//testStruct6のメンバのようにtestStruct5のメンバにアクセスできる
	s.x = x
	s.y = y
}

/*
  実行結果
  -------
  5 10
  -------
*/
