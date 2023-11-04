package log

import (
	"fmt"
	"log"
)

/*
	ログのフォーマットを指定する(SetFlags)

	func SetFlags(flag int)
	  param:
	    flag: ログのフォーマット用定数。論理和で複数組み合わせ可能
		      デフォルトは LstdFlags
	      Ldate         = 1 << iota
          Ltime
          Lmicroseconds
          Llongfile
          Lshortfile
          LUTC
          LstdFlags     = Ldate | Ltime
	  return:
	    なし
*/
func LogSample003() {
	fmt.Println("log_sample_003")

	log.Println("log1")

	//日時のみに変更
	log.SetFlags(log.Ldate)
	log.Println("log2")
}

/*
  実行結果
  -------
  2023/11/04 21:08:28 log1
  2023/11/04 log2
  -------
*/
