package json

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

/*
	JSONファイルをGoオブジェクト(構造体)に読み込む(NewDecoder、Decode)
	func json.NewDecoder(r io.Reader) *json.Decoder
	  param:
	    r            : 読み出すos.Fileオブジェクト
	  return:
	    *json.Decoder: JSONデコーダ

	func (*json.Decoder).Decode(v any) error
	  param:
	    v     : 読み込んだデータを格納するオブジェクト(構造体)
	  return:
	    error : エラー
*/
func JsonSample004() {
	fmt.Println("json_sample_004")

	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	// JSONファイルオープン
	file, err := os.Open("./data/json/person.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// JSONファイルの内容をPerson構造体データとして読み出す
	var p Person
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&p); err != nil {
		log.Fatal(err)
	}

	fmt.Println(p)
}
