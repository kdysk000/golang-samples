package os

import (
	"fmt"
	"log"
	"os"
)

/*
	カレントディレクトのパスを取得(Getwd)
	func Getwd() (dir string, err error)
		param:
		  なし
		return:
		  dir  : ディレクトリパス
		  error: エラー
*/
func OsSample001() {
	fmt.Println("os_sample_001")

	// カレントディレクトのパスを取得
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("current dir:", dir)
}
