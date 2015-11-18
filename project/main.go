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
	fmt.Println("Adder: sum ->", sum, "cout ->", cout)

}
