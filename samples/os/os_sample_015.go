package os

import (
	"fmt"
	"log"
	"os"
)

/*
    シンボリックリンクの作成(Symlink)

	func Symlink(oldname, newname string) error
	  param:
	    oldname: ファイルパス
	    newname: 作成するシンボリックリンクのパス
	  return:
	    error  : エラー

	  注：
	    oldname は、newname から見た相対パスか絶対パスである必要がある
*/
func OsSample015() {
	fmt.Println("os_sample_015")

	//test.txt と同じディレクトリにシンボリックリンクを作成
	err := os.Symlink("./test.txt", "./data/os/link.txt")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Println("symlink create success.")
}
