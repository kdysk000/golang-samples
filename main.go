package main

import (
	"fmt"
	"os"
	"strconv"

	"go-samples/funclist"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Invalid argument. (args count: %d)\n", len(os.Args))
		os.Exit(1)
	}

	switch os.Args[1] {
	case "csv":
	case "filepath":
	case "json":
	case "os":
	case "strings":
	default:
		fmt.Printf("Invalid argument. (args[package]: %s)\n", os.Args[1])
		os.Exit(1)
	}

	no, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("Invalid argument. (args[test no]: %s)\n", os.Args[2])
		os.Exit(1)
	}

	fl := funclist.GetFunc(os.Args[1])
	if no > len(fl) {
		fmt.Printf("Invalid argument. (Max test no is %d)\n", len(fl))
		os.Exit(1)
	}

	fl[no-1]()
}
