package csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strings"
)

/*
	CSV形式のテキストを読み出す(NewReader、Read)

	func NewReader(r io.Reader) *csv.Reader
	  param:
	    r          : ioリーダ(CSVテキストを読みだす場合はstrings.Readerを渡す)
	  return:
	    *csv.Reader: csvリーダ

	func (r *csv.Reader) Read() (record []string, err error)
	  param:
	    なし
	  return:
	    record: CSVから読みだした1レコード分のデータ
	    err   : エラー
*/
func CsvSample002() {
	fmt.Println("csv_sample_002")

text := `name,age,address
hoge,20,tokyo
fuga,30,osaka
`
	reader := csv.NewReader(strings.NewReader(text))
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
