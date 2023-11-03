package os

import (
	"fmt"
	"log"
	"os"
)

/*
	シンボリックリンクのリンク先を読み込む(Readlink)

	func Readlink(name string) (string, error)
	  param:
	    name  : ファイル名
	  return:
	    string: リンク先のパス
	    error : エラー

	  注：
	    戻り値のリンク先のパスは、相対パスの場合シンボリックリンクからのパスになる
*/
func OsSample016() {
	fmt.Println("os_sample_016")

	if ret := IsExists("./data/os/link.txt"); !ret {
		err := os.Symlink("./test.txt", "./data/os/link.txt")
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}

	path, err := os.Readlink("./data/os/link.txt")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Println(path)
}

/*
  実行結果
  -------
  ./test.txt
  -------
*/