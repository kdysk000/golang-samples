package filepath

import (
	"fmt"
	"path/filepath"
)

/*
	パスをディレクトリ名とファイル名に分割する(Split)
	func Split(path string) (dir, file string)
		param:
		  path: パス
		return:
		  dir : ディレクトリ名
		  file: ファイル名
*/
func FilepathSample009() {
	fmt.Println("filepath_sample_009")

	dir, file := filepath.Split("/tmp/dir/test.go")
	fmt.Println(dir)
	fmt.Println(file)
}

/*
  実行結果
  -------
  /tmp/dir/
  test.go
  -------
*/
