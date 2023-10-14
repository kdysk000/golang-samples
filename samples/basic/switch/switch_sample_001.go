package myswitch

import (
	"fmt"
)

/*
	基本的なswitch文

	他の言語にあるような break 文は不要
	逆に、次の条件の処理も実行するには fallthrough を用いる
*/
func SwitchSample001() {
	fmt.Println("switch_sample_001")

	mode := "running"

	switch mode {
	case "running":
		fmt.Println("実行中")
	case "stop":
		fmt.Println("停止中")
	default:
		fmt.Println("不明")
	}
}

/*
  実行結果
  -------
  実行中
  -------
*/
