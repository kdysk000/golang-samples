package basic

import (
	"fmt"
)

/*
	即時関数

	定義と同時に関数を実行する
*/
func BasicSample023() {
	fmt.Println("basic_sample_023")

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
