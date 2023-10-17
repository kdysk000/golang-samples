package os

import (
	"fmt"
	"log"
	"os"
)

/*
    ファイルの書き込み(Write、WriteAt)

	func (f *File) Write(b []byte) (n int, err error)
	  概要:
	    ファイルの先頭から書き込み
	  param:
	    b    : ファイルに書き込むデータを格納したbyte型のスライス
	  return:
	    n    : 書きこんだバイト数
	    error: エラー

	func (f *File) WriteAt(b []byte, off int64) (n int, err error)
	  概要:
	    指定した位置から書き込み
	    指定した位置がファイルサイズを超えている場合は、元のファイル内容の直後から書き込みが行われる
	  param:
	    b    : ファイルに書き込むデータを格納したbyte型のスライス
	    off  : 書き込みを開始するバイト位置(0が先頭)
	  return:
	    n    : 書きこんだバイト数
	    error: エラー
*/
func OsSample010() {
	fmt.Println("os_sample_010")

	// 書き込みの場合はCreate()でファイルを開く
	f, err := os.Create("data/os/test.txt")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer f.Close()

	str1 := "123456789\n"
	data1 := []byte(str1)
	count1, err := f.Write(data1)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Printf("write %d bytes:\n", count1)

	str2 := "ABCDEFGHI\n"
	data2 := []byte(str2)
	count2, err := f.WriteAt(data2, 10)  // 位置10(11バイト目)から書き込む
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Printf("write at %d bytes:\n", count2)
}
