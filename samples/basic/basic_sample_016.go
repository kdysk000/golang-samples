package basic

import (
	"fmt"
)

/*
	switch文のcaseで条件判定

	case 分で条件を記述することが可能
*/
func BasicSample016() {
	fmt.Println("basic_sample_016")

	x := 1
	y := 2

	switch {
	case x > y:
		fmt.Println("Big")
	case x < y:
		fmt.Println("Small")
	default:
		fmt.Println("Equal")
	}
}

/*
  実行結果
  -------
  small
  -------
*/
