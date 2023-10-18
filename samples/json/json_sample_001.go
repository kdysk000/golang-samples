package json

import (
	"encoding/json"
	"fmt"
	"log"
)

/*
	Goオブジェクト(構造体)をJSON形式の文字列データに変換(Marshal)

	func Marshal(v any) ([]byte, error)
	  param:
	    v     : 変換するオブジェクト
	  return:
	    []byte: JSON文字列データ
	    error : エラー
*/
func JsonSample001() {
	fmt.Println("json_sample_001")

	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	p := Person{"hogehoge", 20}
	
	bytes, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bytes))
}
