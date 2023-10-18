package filepath

import (
	"fmt"
	"log"
	"path/filepath"
)

/*
	指定したpatternに一致するパス名をスライスで返す(Glob)

	func Glob(pattern string) (matches []string, err error)
	  param:
	    pattern: 検索したいパターン
	  return:
	    matches: パターンにマッチしたパスのスライス
	    err    :
*/
func FilepathSample006() {
	fmt.Println("filepath_sample_006")

	files, err := filepath.Glob("./data/filepath/*.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(files)
}

/*
  実行結果
  -------
  [data/filepath/tmp1.txt data/filepath/tmp2.txt data/filepath/tmp3.txt]
  -------
*/
