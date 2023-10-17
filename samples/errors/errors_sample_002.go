package errors

import (
	"errors"
	"fmt"
)

/*
	エラーのアンラップ(Unwrap)

	「ラッピングされたエラー」から「ラッピング前のエラー」を取り出す
	※ エラーはfmt.Errorf()に”%w”フォーマットを適用する事によってラッピングできる

	func Unwrap(err error) error
	  param:
	    err: Wrapされたエラー
	  return:
	    Wrap前のエラー
*/
func ErrorsSample002() {
	fmt.Println("errors_sample_002")

	err := errors.New("err1")
	WrappedErr := fmt.Errorf("Wrapped %w", err)

	fmt.Println(err)
	fmt.Println(WrappedErr)
}

/*
  実行結果
  -------
  err1
  Wrapped err1
  -------
*/
