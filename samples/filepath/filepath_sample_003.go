package filepath

import (
	"fmt"
	"path/filepath"
)

/*
	パス名の最後の要素を除いたものを返す(Dir)
	func Dir(path string) string
		param:
		  path パス
		return:
		  string: パスの最後の要素を除いたパス
	注：
	  与えたパスが"/"だけの場合は"/"を返す
	  パスが空の場合は"."を返す
*/
func FilepathSample003() {
	fmt.Println("filepath_sample_003")

	fmt.Println(filepath.Dir("/tmp/dir/temp.txt"))
	fmt.Println(filepath.Dir("/tmp/dir/"))
	fmt.Println(filepath.Dir("/tmp/dir"))
	fmt.Println(filepath.Dir("/"))
	fmt.Println(filepath.Dir(""))
}

/*
  実行結果
  -------
  /tmp/dir
  /tmp/dir
  /tmp
  /
  .
  -------
*/
