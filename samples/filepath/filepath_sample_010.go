package filepath

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

/*
	指定したパス以下のディレクトリを再帰的に探索する(Walk)

	func Walk(root string, fn WalkFunc) error
	  概要:
	    第二引数にWalkFuncを指定するとファイルやディレクトリが見つかるたびにWalkFuncが実行される
	  param:
	    root : 探索するディレクトリ
	    fn   : ファイルやディレクトリが見つかった際に実行する関数
	  return:
	    error: 連結後のパス
*/
func FilepathSample010() {
	fmt.Println("filepath_sample_010")

	err := filepath.Walk("./data/filepath/walk",
		func(path string, info os.FileInfo, err error) error {
			if filepath.Ext(path) == ".txt" {
				fmt.Println(path)
			}
			return nil
		})
	if err != nil {
		log.Fatal(err)
	}
}

/*
  実行結果
  -------
  data/filepath/walk/dir/test3.txt
  data/filepath/walk/test1.txt
  data/filepath/walk/test2.txt  
  -------
*/
