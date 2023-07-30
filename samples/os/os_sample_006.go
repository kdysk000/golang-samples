package os

import (
	"fmt"
	"log"
	"os"
)

/* 
    ファイルを開く(Open)、ファイルを閉じる(Close)
	func Open(name string) (*File, error)
	  param:
	    name: ファイル名
	  return:
	    *File: ファイルオブジェクト
		error: エラー
	  注：
	    Open()は読み込み専用、書き込みする場合はCreate()を使う

	func (f *File) Close() error
	  param:
	    なし
	  return:
		error: エラー
*/
func OsSample006() {
	fmt.Println("os_sample_006")

	f, err := os.Open("data/os/test.txt")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer f.Close()

	fmt.Println("file open success.")
	fmt.Println(f.Name)
}
