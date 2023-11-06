package log

import (
	"fmt"
	"log"
	"os"
)

/*
	ログの出力先を変更する(SetOutput)

	func SetOutput(w io.Writer)
	  param:
	    w: 出力先のI/O Writer
	  return:
	    なし
*/
func LogSample002() {
	fmt.Println("log_sample_002")

	f, err := os.Create("./data/log/log.txt")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	//ログの出力先をファイルに変更
	log.SetOutput(f)
	log.Println("log1")

	//ログの出力先を標準出力に変更
	log.SetOutput(os.Stdout)
	log.Println("log2")
}

/*
  実行結果
  -------
  2023/11/04 20:59:24 log2
  -------
*/
