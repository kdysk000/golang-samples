package os

import (
	"fmt"
	"log"
	"os"
)

/*
	ディレクトの作成(Mkdir)

	func Mkdir(name string, perm FileMode) error
	  param:
	    name : 作成するディレクトリ名を含んだパス
	    perm : パーミッション
	  return:
	    error: エラー
*/
func OsSample003() {
	fmt.Println("os_sample_003")

	// ディレクトリ作成
	err := os.Mkdir("data/os/testdir1", 0755)
	if err != nil {
		log.Fatal(err)
	}
}
