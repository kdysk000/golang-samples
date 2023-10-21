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
func OsSample011() {
	fmt.Println("os_sample_011")

	bytes, err := os.ReadFile("data/os/test.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bytes))
}

/*
  実行結果
  -------
  test

  -------
*/