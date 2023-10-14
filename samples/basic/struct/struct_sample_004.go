package mystruct

import (
	"fmt"
)

/*
	構造体の値レシーバとポインタレシーバ

	値レシーバ：
	　メソッドが呼び出される際に、構造体のコピーが渡される。
	　これにより、メソッド内で構造体のデータを変更しても元の構造体のデータは更新されない。

	ポインタレシーバ：
	　メソッドが呼び出される際に、構造体のアドレス（参照）が渡される。
	　これにより、メソッド内で構造体のデータを直接変更することができる。
	　また、メソッドの呼び出しに伴って構造体がコピーされないため、メモリ使用量が削減される。

*/
type testStruct4 struct {
	x int
	y int
}

func StructSample004() {
	fmt.Println("struct_sample_004")

	st := testStruct4{5,10}

	//値レシーバのメソッドなので元の構造体は更新されない
	st.addNum1(1)
	fmt.Println(st.getNum())

	//ポインタレシーバのメソッドなので元の構造体が更新される
	st.addNum2(1)
	fmt.Println(st.getNum())
}

//値レシーバのgetter
func (s testStruct4) getNum() (int, int) {
	return s.x, s.y
}

//値レシーバのsetter
func (s testStruct4) addNum1(i int) {
	s.x += i
	s.y += i
}

//ポインタレシーバのsetter
func (s *testStruct4) addNum2(i int) {
	s.x += i
	s.y += i
}

/*
  実行結果
  -------
  5 10
  6 11
  -------
*/
