package myinterface

import (
	"fmt"
	"time"
)

/*
	ゴルーチンの待ち合わせ(channel)

	チャネルとは
	  ・Go言語に組み込まれた変数の種類の1つで、ゴルーチン間でデータの送受信を行うために使用する
	  ・チャネルはmake()で生成することができる
	  ・使い終わったらclose()でチャネルをクローズする
	  ・チャネルにメッセージを送信するには、 「 チャネル変数 <- 値 」 と書く
	  ・チャネルからメッセージを受信するには、 「 変数 := <- チャネル変数 」 と書く

	書式
	  ch := make(chan 型)
	  close(ch)
*/
func funcB(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "END"  // チャネルにメッセージを送信する
}

func GoroutineSample002() {
	fmt.Println("goroutine_sample_002")

	// チャネル生成
	ch := make(chan string)
	defer close(ch)

	// チャネルをゴルーチンに渡して実行
    go funcB(ch)

	// チャネルからメッセージを受信する(受信するまで待機)
    msg := <- ch 

    fmt.Println(msg)
}

/*
  実行結果
  -------
  END
  -------
*/
