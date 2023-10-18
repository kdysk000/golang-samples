package csv

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

/*
	CSVファイルに書き出す(NewWriter、Writer、Flash)

	func NewWriter(w io.Writer) *csv.Writer
	  param:
	    w          : ioライタ(ファイルに書き出す場合はos.Fileを渡す)
	  return:
	    *csv.Writer: csvライタ

	func (w *csv.Writer) Write(record []string) error
	  param:
	    record: CSVに書き出す1レコード分のデータ
	  return:
	    err   : エラー
*/
func CsvSample004() {
	fmt.Println("csv_sample_004")

	// 出力用ファイルをオープン
	file, err := os.Create("./data/csv/write_test.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	// 内部バッファのフラッシュ
	defer writer.Flush() 

	writer.Write([]string{"AAA", "BBB", "CCC"})
	writer.Write([]string{"100", "200", "300"})
	writer.Write([]string{"XXX", "YYY", "ZZZ"})
}

/*
  実行結果
  -------
  [name age address]
  [hoge 20 tokyo]
  [fuga 30 osaka]
  -------
*/
