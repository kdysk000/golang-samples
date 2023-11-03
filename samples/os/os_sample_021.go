package os

import (
	"fmt"
	"os"
)

/*
	環境変数の削除(Unsetenv)

	func Unsetenv(key string) error
	  param:
	    key　: 環境変数名
	  return:
	    error: エラー
*/
func OsSample021() {
	fmt.Println("os_sample_021")

	os.Setenv("TEST", "test")
	fmt.Println("Environment(TEST):", os.Getenv("TEST"))

	os.Unsetenv("TEST")
	fmt.Println("Environment(TEST):", os.Getenv("TEST"))
}

/*
  実行結果
  -------
  Environment(TEST): test
  Environment(TEST):
  -------
*/
