package filepath

import (
	"fmt"
	"path/filepath"
)

/*
	linkを指定した場合、リンク先のパスを返す(EvalSymlinks)

	func EvalSymlinks(path string) (string, error)
	  param:
	    path  : linkパス
	  return:
	    string: リンク先のパス
	    error :
*/
func FilepathSample004() {
	fmt.Println("filepath_sample_004")

	fmt.Println(filepath.EvalSymlinks("./data/filepath/004"))
}

/*
  実行結果
  -------
  samples/filepath/filepath_sample_004.go <nil>
  -------
*/
