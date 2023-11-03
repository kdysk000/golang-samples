package os

import (
	"fmt"
	"os"
)

/*
	環境変数の設定(Setenv)

	func Setenv(key, value string) error
	  param:
	    key　 : 環境変数名
	    value : 環境変数の値
	  return:
	    error: エラー
*/
func OsSample020() {
	fmt.Println("os_sample_020")

	os.Setenv("TEST", "test")
	fmt.Println("Environment(TEST):", os.Getenv("TEST"))
}

/*
  実行結果
  -------
  Environment(TEST): test
  -------
*/
