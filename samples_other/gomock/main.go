package main

import (
	"fmt"
)

type NumIf interface {
	GetNum() int
}

type Number struct {
	num int
}

func (n *Number) GetNum() int {
	return n.num
}

type ExeCalc struct {
	n NumIf
}

func (ec *ExeCalc) Plus(num int) int {
	return ec.n.GetNum() + num
}

func (ec *ExeCalc) Minus(num int) int {
	return ec.n.GetNum() - num
}

func main() {
	number := &Number{num:10}
	calc := &ExeCalc{n:number}
	fmt.Println(calc.Plus(10))
	fmt.Println(calc.Minus(10))
}