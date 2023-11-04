package log

import (
	"fmt"
	"log"
)

/*
	ログのプリフィックスを指定する(SetPrefix)

	func SetPrefix(prefix string)
	  param:
	    prefix: ログの先頭に追加するPrefix(文字列)
	  return:
	    なし
*/
func LogSample004() {
	fmt.Println("log_sample_004")

	log.Println("log1")

	//日時のみに変更
	log.SetPrefix("[LOG]")
	log.Println("log2")
}

/*
  実行結果
  -------
  2023/11/04 21:11:45 log1
  [LOG]2023/11/04 21:11:45 log2
  -------
*/
