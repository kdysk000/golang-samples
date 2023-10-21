package os

import (
	"fmt"
	"os"
)

/*
	環境変数の設定(Setenv)

	func Setenv(key, value string) error
*/
func OsSample005() {
	fmt.Println("os_sample_005")

	os.Setenv("TEST", "test")
	fmt.Println("Environment(TEST):", os.Getenv("TEST"))
}

/*
  実行結果
  -------
  Environment(TEST): test
  -------
*/
