package csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

/*
	CSV形式のファイルを読み出す(NewReader、Read)
	func NewReader(r io.Reader) *csv.Reader
	  param:
	    r          : ioリーダ(ファイルを読みだす場合はos.Fileを渡す)
	  return:
	    *csv.Reader: csvリーダ

	func (r *csv.Reader) Read() (record []string, err error)
	  param:
	    なし
	  return:
	    record: CSVから読みだした1レコード分のデータ
	    err   : エラー
*/
func CsvSample001() {
	fmt.Println("csv_sample_001")

	// CSVファイルオープン
	file, err := os.Open("./data/csv/test.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	for {
		// CSVを1レコード読み込み
		record, err := reader.Read()
		// 最後まで読み出したら抜ける
		if err == io.EOF {
			break 
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(record)
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
