package log

import (
	"fmt"
	"log"
)

/*
	ログの出力(Print/Println/Printf)

	func Print(v ...any)
	func Println(v ...any)
	func Printf(format string, v ...any)

	・PrintとPrintlnの違い
	　・Print
	　　・引数の文字列の最後に改行の有無に関わらずメッセージは改行される（2度改行されて空行が表示されたりはしない）
	　　・複数の引数を与えると fmt.Print()のように引数の自然な文字列表現をそのままつなげて出力する
	　・Println
	　　・最後に改行があれば改行を入れる(2度改行されることになる)
	　　・複数の引数がある場合はそれぞれの自然な文字列表現をスペースで連結して出力する
*/
func LogSample001() {
	fmt.Println("log_sample_001")

	log.Print("log1\n")
	log.Println("log2\n")

	log.Print("Print:", "log1")
	log.Println("Println:", "log2")

	log.Printf("Printf: %s", "log3")
}

/*
  実行結果
  -------
  2023/10/07 18:51:50 log1
  2023/10/07 18:51:50 log2

  2023/10/07 18:51:50 Print:log1
  2023/10/07 18:51:50 Println: log2
  2023/10/07 18:51:50 Printf: log3
  -------
*/
