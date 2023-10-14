package basic

import (
	"fmt"
)

/*
	構造体メンバのpublicとprivate

	大文字で始まるものはパッケージ外からアクセス可能(Public)であり、
	小文字で始まるものはパッケージ外からアクセス不可(private)となる。
*/
type testStruct2 struct {
	a int
	B int
}

func BasicSample027() {
	fmt.Println("basic_sample_027")

	st := testStruct2{1, 2}

	fmt.Println(st.a)  //Privateなメンバだが同じパッケージ内なのでアクセス可能
	fmt.Println(st.B)
}

/*
  実行結果
  -------
  1
  2
  -------
*/
