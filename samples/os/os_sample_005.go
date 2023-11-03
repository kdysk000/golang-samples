package os

import (
	"fmt"
	"log"
	"os"
)

/*
    テンポラリディレクトリの作成(MkdirTemp)

	func MkdirTemp(dir string, pattern string) (string, error)
	  param:
	    dir    : ディレクトリの作成場所
	    pattern: 作成するディレクトリ名の先頭部分
	  return:
	    string : 作成したディレクトを含めたパス
	    error  : エラー
*/
func OsSample005() {
	fmt.Println("os_sample_005")

	// テンポラリディレクトリ作成
	dir, err := os.MkdirTemp("data/os", "sample")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("dir:", dir)
}

/*
  実行結果
  -------
  dir: data/os/sample2547357272
  -------
*/