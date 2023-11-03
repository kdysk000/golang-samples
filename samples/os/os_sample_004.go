package os

import (
	"fmt"
	"log"
	"os"
)

/*
	ディレクトの一括作成(MkdirAll)

	func MkdirAll(name string, perm FileMode) error
	  param:
	    name : 作成するディレクトリ名を含んだパス
	    perm : パーミッション
	  return:
	    error: エラー

	foo/bar/bazのように階層の深いディレクトリを一括して作成する
*/
func OsSample004() {
	fmt.Println("os_sample_004")

	// ディレクトリ作成
	err := os.MkdirAll("data/os/testdir2/subdir1", 0755)
	if err != nil {
		log.Fatal(err)
	}
}
