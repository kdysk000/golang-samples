package myfunc

import (
	"fmt"
)

/*
	即時関数

	定義と同時に関数を実行する
*/
func FuncSample009() {
	fmt.Println("func_sample_009")

	greeting := func(country string) string {
		switch country {
		case "Japan":
			return "こんにちは"
		case "America":
			return "Hello"
		default:
			return "xxxxx"
		}
	}("Japan")

	fmt.Println(greeting)
}

/*
  実行結果
  -------
  こんにちは
  -------
*/
