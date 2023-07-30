package os

import (
	"fmt"
	"os"
)

/* 
    環境変数の取得(Getenv)
	func Getenv(key string) string
	  param:
	    key　 : 環境変数名
	  return:
		string: 環境変数の値
*/
func OsSample003() {
	fmt.Println("os_sample_003")

	// カレントディレクトのパスを取得
	fmt.Println("Environment(HOME):", os.Getenv("HOME"))
}