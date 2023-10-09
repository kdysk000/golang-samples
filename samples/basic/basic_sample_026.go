package basic

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

func BasicSample026() {
	fmt.Println("basic_sample_026")

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
