package myinterface

import (
	"fmt"
	"time"
)

/*
	チャネルが値を受信したかクローズされたかの区別

	msg := <- ch で受信待ち中にチャネルがcloseされた場合は、ゼロ値が受信される
	値を受信したのかcloseによってゼロ値になったのかを区別するには、左辺の2つめの変数でbool値を受け取ることができる

	書式
	  msg, ok := <- ch
	    値を受信した場合           ：ok が true
		チャネルがクローズされた場合：ok が false
*/
func funcC(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "END"  // チャネルにメッセージを送信する
}

func GoroutineSample003() {
	fmt.Println("goroutine_sample_003")

	// チャネル生成
	ch := make(chan string)
	defer close(ch)

	// チャネルをゴルーチンに渡して実行
	go funcC(ch)

	// チャネルからメッセージを受信する
	msg, ok := <- ch
	if ok {
		fmt.Println(msg)
	} else {
		fmt.Println("closed")
	}    
}

/*
  実行結果
  -------
  END
  -------
*/
