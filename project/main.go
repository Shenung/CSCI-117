package main

import (
	"fmt"

	"./components"
)

func typeConv(a byte) int {
	if a == 48 {
		return 0
	}
	return 1
}

func main() {
	var a, b = make([]byte, 16), make([]byte, 16)

	fmt.Println("Enter 16bit binary [A]: ")
	fmt.Scanln(&a)
	fmt.Println("Enter 16bit binary [B]: ")
	fmt.Scanln(&b)

	fmt.Println(a)
	fmt.Println(b)

	for i := 15; i >= 0; i-- {
		if a[i] == 48 || a[i] == 49 && b[i] == 48 || b[i] == 49 {
			fmt.Println("from a->", typeConv(a[i]))
			fmt.Println("from b->", typeConv(b[i]))
		}
	}

	fmt.Println("And gate: ", components.AndGate(1, 1))
	fmt.Println("Or gate: ", components.OrGate(1, 1))
	fmt.Println("Xor gate: ", components.XorGate(1, 1))
	sum, cout := components.Adder(1, 1, 1)
	fmt.Println("Adder: sum ->", sum, " cout ->", cout)
	result := components.Mux4x1(components.AndGate(1, 1), components.OrGate(1, 1), sum, 0, 00)
	fmt.Println("Mux4x1: ", result)
	fresult, cout := components.ALU1bit(1, 1, 1, 00)
	fmt.Println("ALU1bit: result->", fresult, " cout->", cout)
}
