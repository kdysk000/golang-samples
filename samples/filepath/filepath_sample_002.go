package filepath

import (
	"fmt"
	"path/filepath"
)

/*
	パスの最後の要素を返す(Base)
	func Base(path string) string
		param:
		  path パス
		return:
		  string: パスの最後の要素
	注：
	  与えたパスが"/"だけの場合は"/"を返す
	  パスが空の場合は"."を返す
*/
func FilepathSample002() {
	fmt.Println("filepath_sample_002")

	fmt.Println(filepath.Base("/tmp/dir/test.go"))
	fmt.Println(filepath.Base("/"))
	fmt.Println(filepath.Base(""))
}

/*
  実行結果
  -------
  test.go
  /
  .
  -------
*/