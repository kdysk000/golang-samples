package mystruct

import (
	"fmt"
)

/*
	構造体の比較

	・基本型のみのフィールドを持つ構造体は == で比較可能
	・フィールドにスライスなどがある場合は == ではエラーとなる
	・reflectパッケージのDeepEqual()などを使って比較する必要がある
*/
type testStruct6 struct {
	x int
	y int
}

func StructSample006() {
	fmt.Println("struct_sample_006")

	st1 := testStruct6{1, 2}
	st2 := testStruct6{1, 2}
	st3 := testStruct6{1, 3}

	fmt.Println(st1 == st2)
	fmt.Println(st1 == st3)
}

/*
  実行結果
  -------
  true
  false
  -------
*/
