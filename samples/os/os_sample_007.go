package os

import (
	"fmt"
	"log"
	"os"
)

/*
	書き込み用にファイルを開く(Create)

	func Create(name string) (*File, error)
	  param:
	    name: ファイル名
	  return:
	    *File: ファイルオブジェクト
	    error: エラー
	  注：
	    Create()関数は、ファイルが既に存在している場合、中身を空にして開く。
	    ファイルが存在していなかった場合はファイルを作成する。
*/
func OsSample007() {
	fmt.Println("os_sample_007")

	fp, err := os.Create("data/os/test007.txt")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer fp.Close()

	fmt.Println("file create success.")
}

/*
  実行結果
  -------
  file create success.
  -------
*/
