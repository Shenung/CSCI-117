package main

import (
	"fmt"

	"./components"
)

func main() {
	fmt.Println("Enter 16bit binary [A]: ")

	fmt.Println("Enter 16bit binary [B]: ")

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
