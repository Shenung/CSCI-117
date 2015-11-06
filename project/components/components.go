package componenets

import "fmt"

//Basic gates And, Or, Not, and Xor

//AndGate returns bitwise And
func AndGate(a, b int) int {
	return a & b
}

//OrGate returns bitwise Or
func OrGate(a, b int) int {
	return a | b
}

//NotGate returns 0 or 1 depending on the input
func NotGate(a int) int {
	if a == 1 {
		return 0
	}
	return 1
}

//XorGate takes in 2 inputs and returns an int
func XorGate(a, b int) int {
	temp1, temp2 := NotGate(a), NotGate(b)
	a1temp, a2temp := AndGate(a, temp2), AndGate(b, temp1)
	return OrGate(a1temp, a2temp)
}

//Adder adds to signals together
func Adder(a, b, c int) (int, int) {
	temp1 := XorGate(a, b)
	sum := XorGate(temp1, c)

	tempA1 := AndGate(a, b)
	tempA2 := AndGate(temp1, c)
	cout := OrGate(tempA1, tempA2)
	return sum, cout
}

//Mux4x1 takes in 4 signals and uses and op code to figure out which result to spit out
func Mux4x1(a, b, c, d, op int) int {
	switch {
	case op == 00:
		temp1 := AndGate(a, 1)
		temp2 := OrGate(temp1, 0)
		return OrGate(temp2, 0)
	case op == 01:
		temp1 := AndGate(b, 1)
		temp2 := OrGate(temp1, 0)
		return OrGate(temp2, 0)
	case op == 10:
		temp1 := AndGate(c, 1)
		temp2 := OrGate(temp1, 0)
		return OrGate(temp2, 0)
	case op == 11:
		temp1 := AndGate(d, 1)
		temp2 := OrGate(temp1, 0)
		return OrGate(temp2, 0)
	default:
		fmt.Println("Incorrect OP code")
	}
	return 0
}
