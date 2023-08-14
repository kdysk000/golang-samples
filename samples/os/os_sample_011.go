package os

import (
	"fmt"
	"log"
	"os"
)

/*
    ファイルの書き込み(WriteFile)
	func WriteFile(name string, data []byte, perm FileMode) error
		param:
		  name : ファイルパス
		  data : ファイルに書き込むデータを格納したbyte型のスライス
		  perm : パーミッション
		return:
		  error: エラー
		注：
		  ファイルが既に存在している場合、中身を空にしてデータを書き込む。
		  ファイルが存在していなかった場合はファイルを作成する。
*/
func OsSample011() {
	fmt.Println("os_sample_011")

	str := "write this file by Golang!"
	data := []byte(str)
	err := os.WriteFile("data/os/test.txt", data, 0755)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Write file success.")
}
