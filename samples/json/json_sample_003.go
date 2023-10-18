package json

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

/*
	Goオブジェクト(構造体)をJSONファイルに出力する(NewEncoder、Encode)

	func json.NewEncoder(w io.Writer) *json.Encoder
	  param:
	    w            : ファイルに書き出すos.Fileオブジェクト
	  return:
	    *json.Encoder: JSONエンコーダ

	func (*json.Encoder).Encode(v any) error
	  param:
	    v     : 書き込むオブジェクト(構造体)
	  return:
	    error : エラー
*/
func JsonSample003() {
	fmt.Println("json_sample_003")

	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	// 書き込むデータ
	p := Person{"hogehoge", 30}

	// JSONファイルを作成
	file, err := os.Create("./data/json/person.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// JSONファイルに書き込み
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(p); err != nil {
		log.Fatal(err)
	}
}
