package os

import (
	"fmt"
	"os"
)

/*
	環境変数の取得(Environ)

	func Environ() []string
	  return:
	    string: 環境変数

	「key=value」の形式で全ての環境変数を返す
*/
func OsSample019() {
	fmt.Println("os_sample_019")

	for _, env := range os.Environ() {
		fmt.Println(env)
	}
}