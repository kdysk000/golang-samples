package errors

import (
	"errors"
	"fmt"
)

/*
	エラーの判別(Is)

	「特定のエラーであるのか」、「特定のエラーをWrapしたエラーであるのか」を判別する
	第一引数にラップしたエラーを指定すると、順番にアンラップして確認する

	func Is(err error, target error) bool
	  param:
	    err   : エラー
	    target:
	  return:
	    bool
*/
func ErrorsSample003() {
	fmt.Println("errors_sample_003")

	err1 := errors.New("err1")
	err2 := errors.New("err2")
	err3 := errors.New("err1")  //エラー文字列はerr1と同じ
	WrappedErr1 := fmt.Errorf("Wrapped %w", err1)  //err1をWrap

	fmt.Println(errors.Is(err1, err2))
	fmt.Println(errors.Is(err1, err3))         //中身は同じだが別エラー
	fmt.Println(errors.Is(err1, WrappedErr1))  //Wrapされたほうが第2引数なので等価にならない
	fmt.Println(errors.Is(WrappedErr1, err1))  //Wrapされたほうが第1引数なので順番にUnwrapされて評価され等価となる
}

/*
  実行結果
  -------
  false
  false
  false
  true
  -------
*/
