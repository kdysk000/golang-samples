package errors

import (
	"errors"
	"fmt"
)

/*
	エラーの定義(New)

	func New(text string) error
	  param:
	    text: エラー文字列
	  return:
	    error
*/
func ErrorsSample001() {
	fmt.Println("errors_sample_001")

	err := errors.New("err1")
	fmt.Println(err)
}

/*
  実行結果
  -------
  err1
  -------
*/
