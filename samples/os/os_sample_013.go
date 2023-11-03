package os

import (
	"fmt"
	"log"
	"os"
)

/*
    ファイル名の変更と移動(Rename)

	func Rename(oldpath, newpath string) error
	  param:
	    oldpath: 変更元のファイルパス
	    newpath: 変更後のファイルパス
	  return:
	    error  : エラー

	第1引数と第2引数でファイル名を除くパスが一致すればリネーム、パスが異なれば移動
*/
func OsSample013() {
	fmt.Println("os_sample_013")

	//ファイル作成
	fp, err := os.Create("data/os/before.txt")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fp.Close()

	//リネーム
	err = os.Rename("data/os/before.txt", "data/os/after.txt")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Println("file rename success.")
}
