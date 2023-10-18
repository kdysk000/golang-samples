package filepath

import (
	"fmt"
	"path/filepath"
)

/*
	パス名から拡張子を取得する(Ext)

	func Ext(path string) string
	  param:
	    path  : パス
	  return:
	    string: 取得した拡張子
	注：
	  拡張子がなければ空文字列が返る
*/
func FilepathSample005() {
	fmt.Println("filepath_sample_005")

	fmt.Println(filepath.Ext("/tmp/dir/temp.txt"))
	fmt.Println(filepath.Ext("/tmp/dir/"))
	fmt.Println(filepath.Ext("/tmp/dir"))
}

/*
  実行結果
  -------
  .txt
  
  
  -------
*/
