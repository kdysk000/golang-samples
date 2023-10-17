package filepath

import (
	"fmt"
	"path/filepath"
)

/*
	パスが絶対パスかどうかを判定する(IsAbs)
	func IsAbs(path string) bool
	  param:
	    path: パス
	  return:
	    bool: 絶対パスならtrue、相対パスならfalse
*/
func FilepathSample007() {
	fmt.Println("filepath_sample_007")

	fmt.Println(filepath.IsAbs("/tmp/dir/temp.txt"))
	fmt.Println(filepath.IsAbs("../temp.txt"))
}

/*
  実行結果
  -------
  true
  false
  -------
*/
