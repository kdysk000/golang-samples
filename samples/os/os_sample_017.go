package os

import (
	"fmt"
	"log"
	"os"
)

/*
    ホスト名の取得(Hostname)

	func Hostname() (name string, err error)
	  return:
	    name: ホスト名
	    err : エラー
*/
func OsSample017() {
	fmt.Println("os_sample_017")

	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Println("hostname:", hostname)
}
