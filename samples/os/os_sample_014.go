package os

import (
	"fmt"
	"log"
	"os"
)

/*
	ファイル、ディレクトリの削除(Remove)

	func Remove(name string) error
	  param:
	    name : ファイルパス
	  return:
	    error: エラー
	  注：
	    ディレクトリは空ディレクトリのみ削除可能
*/
func OsSample014() {
	fmt.Println("os_sample_014")

	err := os.Remove("data/os/testdir")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Remove file success.")
}
