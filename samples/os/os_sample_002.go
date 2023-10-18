package os

import (
	"fmt"
	"log"
	"os"
)

/*
	ディレクトの移動(Chdir)

	func Chdir(dir string) error
	  param:
	    dir  : 移動するディレクトリのパス
	  return:
	    error: エラー
*/
func OsSample002() {
	fmt.Println("os_sample_002")

	// カレントディレクトのパスを取得
	before, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// 親ディレクトリに移動
	err = os.Chdir("../")
	if err != nil {
		log.Fatal(err)
	}

	// カレントディレクトのパスを取得
	after, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("before dir:", before)
	fmt.Println("after dir :", after)
}
