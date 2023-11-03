package os

import (
	"fmt"
	"os"
)

/*
	環境変数の存在チェック(LookupEnv)

	func LookupEnv(key string) (string, bool)
	  param:
	    key　 : 環境変数名
	  return:
	    string: 存在すれば環境変数値
	    bool  : 存在チェック結果
*/
func OsSample022() {
	fmt.Println("os_sample_022")

	os.Setenv("TEST", "test")
	val, ok := os.LookupEnv("TEST")
	fmt.Println("ok:", ok, "TEST:", val)

	//存在しない環境変数
	val, ok = os.LookupEnv("HOGEHOGE")
	fmt.Println("ok:", ok, "HOGEHOGE:", val)
}

/*
  実行結果
  -------
  ok: true TEST: test
  ok: false HOGEHOGE:
  -------
*/
