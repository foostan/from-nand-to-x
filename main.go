package main

import (
	"fmt"
	"github.com/foostan/from-nand-to-x/logic_gate"
)

func main() {
	nand := logic_gate.Nand{}
	fmt.Println(nand.Calc(false, false))
	fmt.Println(nand.Calc(true, false))
	fmt.Println(nand.Calc(false, true))
	fmt.Println(nand.Calc(true, true))
}
