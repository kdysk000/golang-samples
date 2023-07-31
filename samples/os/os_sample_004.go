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
func OsSample004() {
	fmt.Println("os_sample_004")

	// ディレクトリ作成
	err := os.Mkdir("data/os/testdir", 0755)
	if err != nil {
		log.Fatal(err)
	}
}
