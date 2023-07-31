package os

import (
	"fmt"
	"log"
	"os"
)

/* 
    ファイルの読み込み(ReadFile)
	func ReadFile(name string) ([]byte, error)
		param:
		  name   : ファイルパス
		return:
		  []byte : ファイルから読み込んだデータを格納するbyte型のスライス
		  error　: エラー
*/
func OsSample009() {
	fmt.Println("os_sample_009")

	bytes, err := os.ReadFile("data/os/test.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bytes))
}
