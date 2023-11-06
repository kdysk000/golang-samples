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
func OsSample012() {
	fmt.Println("os_sample_012")

	// ディレクトリ作成
	err := os.Mkdir("data/os/testdir", 0755)
	if err != nil {
		log.Fatal(err)
	}

	// ディレクトリ削除
	err = os.Remove("data/os/testdir")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Remove directory success.")
}

/*
  実行結果
  -------
  Remove directory success.
  -------
*/
