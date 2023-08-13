package csv

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

/*
	CSV形式のファイルを一括で読み出す(ReadAll)
	func (r *csv.Reader) ReadAll() (records [][]string, err error)
		param:
		  なし
		return:
		  records: CSVから読みだした全レコードデータ
		  err    : エラー
*/
func CsvSample003() {
	fmt.Println("csv_sample_003")

	// CSVファイルオープン
	file, err := os.Open("./data/csv/test.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// records は [][]string 形式のデータ
	for _, fields := range records {
		fmt.Println(fields)
	}
}

/*
  実行結果
  -------
  [name age address]
  [hoge 20 tokyo]
  [fuga 30 osaka]
  -------
*/
