package filepath

import (
	"fmt"
	"log"
	"path/filepath"
)

/*
	相対パスを絶対パスに変換(Abs)
	func Abs(path string) (string, error)
	  param:
	    path  : 相対パス
	  return:
	    string: 絶対パス
	    error : エラー
*/
func FilepathSample001() {
	fmt.Println("filepath_sample_001")

	// 絶対パスに変換
	path, err := filepath.Abs("./main.go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(path)
}
