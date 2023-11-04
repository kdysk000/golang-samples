package log

import (
	"fmt"
	"log"
	"os"
)

/*
	ロガーの生成(New)

	func New(out io.Writer, prefix string, flag int) *Logger
	  param:
	    out    : 出力先
	    prefix : プレフィックス
	    flag   : フォーマット
	  return:
	    *Logger: ロガー
*/
func LogSample005() {
	fmt.Println("log_sample_005")

	logger := log.New(os.Stdout, "[LOG]", log.Ldate)
	logger.Println("log1")
}

/*
  実行結果
  -------
  2023/11/04 21:08:28 log1
  2023/11/04 log2
  -------
*/
