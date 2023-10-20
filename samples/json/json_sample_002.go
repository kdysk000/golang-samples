package json

import (
	"encoding/json"
	"fmt"
	"log"
)

/*
	JSON形式の文字列データ([]byte)をGoのオブジェクトに変換する(Unmarshal)

	func Unmarshal(data []byte, v any) error
	  param:
	    []byte: JSON文字列データ
	    v     : 変換結果を格納するオブジェクト
	  return:
	    error : エラー
*/
func JsonSample002() {
	fmt.Println("json_sample_002")

	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	jsonBytes := []byte(`{"name":"hogehoge","age":30}`)
	var p Person
	if err := json.Unmarshal(jsonBytes, &p); err != nil {
		log.Fatal(err)
	}

	fmt.Println(p.Name)
	fmt.Println(p.Age)
}

/*
  実行結果
  -------
  hogehoge
  30
  -------
*/