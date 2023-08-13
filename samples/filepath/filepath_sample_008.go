package filepath

import (
	"fmt"
	"path/filepath"
)

/*
	パス名を連結する(Join)
	func Join(elem ...string) string
		param:
		  elem     : ベースのパス名
		  ...string: 連結するパス(可変長)
		return:
		  string   : 連結後のパス
*/
func FilepathSample008() {
	fmt.Println("filepath_sample_008")

	fmt.Println(filepath.Join("abc/", "def"))
	fmt.Println(filepath.Join("abc/", "/def"))
	fmt.Println(filepath.Join("abc", "def", "hij"))
}

/*
  実行結果
  -------
  abc/def
  abc/def
  abc/def/hij
  -------
*/
