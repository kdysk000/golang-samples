package os

import (
	"fmt"
	"io"
	"log"
	"os"
)

/*
    ファイルの読み込み(Read、ReadAt)

	func (f *File) Read(b []byte) (n int, err error)
	  概要：
	    先頭からファイルを読み込む
	  param:
	    b    : ファイルから読み込んだデータを格納するbyte型のスライス
	  return:
	    n    : 読みこんだバイト数
	    error: エラー

	func (f *File) ReadAt(b []byte, off int64) (n int, err error)
	  概要：
	    指定した位置からファイルを読み込む
	  param:
	    b    : ファイルから読み込んだデータを格納するbyte型のスライス
	    off  : 読み取りを開始するバイト位置(0が先頭)
	  return:
	    n    : 読みこんだバイト数
	    error: エラー

	  注：
	    Read()は読み込んだバイト数が0のときだけ io.EOF を返すが、
	    ReadAt()はファイル内容が充分になければ読み込んだバイト数が0でなくても io.EOF を返す
*/
func OsSample008() {
	fmt.Println("os_sample_008")

	f, err := os.Open("data/os/test.txt")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer f.Close()

	data1 := make([]byte, 256)
	count1, err := f.Read(data1)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Printf("read %d bytes:\n", count1)
	fmt.Println(string(data1))

	data2 := make([]byte, 256)
	count2, err := f.ReadAt(data2, 2)  // 位置2(3バイト目)から読む込み
	if err == io.EOF {
		fmt.Printf("read at %d bytes:\n", count2)
		fmt.Println(string(data2))
	} else if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}




}
