package myinterface

import (
	"fmt"
	"time"
)

/*
	ブロックせずにチャネルから受信する方法(select)

	msg := <- ch が実行されるとチャネルに値がまだ入っていない場合、受信するまでブロックされる
	ブロックせずに処理を行う場合はselectを使用する

	書式
	  select {
	  case msg, ok := <- ch:
	  	if ok {
      		・・・
      	} else {
      		・・・
      	}
	  default:
	  	・・・
	  }

	・ch に値が入っている場合は case msg := <- ch の 処理が実行される
	・ch に値が入っていない場合は default の処理が実行される

*/
func funcD(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "END"  // チャネルにメッセージを送信する
}

func GoroutineSample004() {
	fmt.Println("goroutine_sample_004")

	// チャネル生成
	ch := make(chan string)
	defer close(ch)

	// チャネルをゴルーチンに渡して実行
    go funcD(ch)

	select {
	case msg, ok := <- ch:
		if ok {
			fmt.Println(msg)
		} else {
			fmt.Println("closed")
		}
	default:
		//このプログラムではここが実行される
		//funcDは1秒sleep後にメッセージを送信するため、先にselectが実行されるため(ブロックされず先に進むため)
		fmt.Println("no message")
	}
}

/*
  実行結果
  -------
  no message
  -------
*/
