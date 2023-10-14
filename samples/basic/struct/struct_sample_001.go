package mystruct

import (
	"fmt"
)

/*
	構造体の定義

	書式
		type 構造体名 struct {
			x int
			y int
		}
*/
type testStruct1 struct {
	x int
	y int
}

func StructSample001() {
	fmt.Println("struct_sample_001")

	st := testStruct1{
		x: 1,
		y: 2,
	}

	fmt.Println(st.x)
	fmt.Println(st.y)
}

/*
  実行結果
  -------
  1
  2
  -------
*/
