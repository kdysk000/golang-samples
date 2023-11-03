package os

import (
	"fmt"
	"os"
)

/*
    ファイルの存在確認(stat、IsNotExist)

	stat() の戻り値の err が nil なら存在すると判定できる

	func Stat(name string) (FileInfo, error)
	  param:
	    name    : ファイルパス
	  return:
	    FileInfo: ファイル情報
	    error   : エラー
*/
func IsExists(filename string) bool {
    _, err := os.Stat(filename)
    return err == nil
}

func OsSample014() {
	fmt.Println("os_sample_014")

	//存在するファイル
    ret := IsExists("./data/os/test.txt")
	fmt.Println("ret:", ret)

	//存在しないファイル
    ret = IsExists("./data/os/unknown.txt")
	fmt.Println("ret:", ret)
}

/*
  実行結果
  -------
  ret: true
  ret: false
  -------
*/