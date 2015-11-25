package components

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

	tempA1 := AndGate(temp1, c)
	tempA2 := AndGate(a, b)
	cout := OrGate(tempA1, tempA2)
	return sum, cout
}

//Mux4x1 takes in 4 signals and uses and op code to figure out which result to spit out
// 00 -> AND
// 01 -> OR
// 10 -> ADD
// 11 -> Dead signal
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

//ALU1bit uses the sub components AND, OR, NOT, XOR, Adder, and Mux4x1
func ALU1bit(a, b, Cin, op int) (int, int) {
	sum, out := Adder(a, b, Cin)
	result := Mux4x1(AndGate(a, b), OrGate(a, b), sum, 0, op)
	return result, out
}

//ALU1bitOF is the last alu to run and checks for overflow
func ALU1bitOF(a, b, Cin, op int) (int, int) {
	sum, out := Adder(a, b, Cin)
	overflow := XorGate(Cin, out)
	result := Mux4x1(AndGate(a, b), OrGate(a, b), sum, 0, op)
	return result, overflow
}

//used to convert the contents inside the slice of bytes
//inside the slice are the ascii of the 1's and 0's so they have to be checked and converted to the usable binary
func typeConv(a byte) int {
	if a == 48 {
		return 0
	} else if a == 49 {
		return 1
	}
	return int(a)
}

func typeConv2(a int) uint8 {
	if a == 0 {
		return 0
	}
	return 1
}

//ALU16bit simulates a 16 bit MIPS processor
func ALU16bit(a, b []byte, op int) ([]uint8, int) {
	a0, out0 := ALU1bit(typeConv(a[15]), typeConv(b[15]), 0, op)
	a1, out1 := ALU1bit(typeConv(a[14]), typeConv(b[14]), out0, op)
	a2, out2 := ALU1bit(typeConv(a[13]), typeConv(b[13]), out1, op)
	a3, out3 := ALU1bit(typeConv(a[12]), typeConv(b[12]), out2, op)
	a4, out4 := ALU1bit(typeConv(a[11]), typeConv(b[11]), out3, op)
	a5, out5 := ALU1bit(typeConv(a[10]), typeConv(b[10]), out4, op)
	a6, out6 := ALU1bit(typeConv(a[9]), typeConv(b[9]), out5, op)
	a7, out7 := ALU1bit(typeConv(a[8]), typeConv(b[8]), out6, op)
	a8, out8 := ALU1bit(typeConv(a[7]), typeConv(b[7]), out7, op)
	a9, out9 := ALU1bit(typeConv(a[6]), typeConv(b[6]), out8, op)
	a10, out10 := ALU1bit(typeConv(a[5]), typeConv(b[5]), out9, op)
	a11, out11 := ALU1bit(typeConv(a[4]), typeConv(b[4]), out10, op)
	a12, out12 := ALU1bit(typeConv(a[3]), typeConv(b[3]), out11, op)
	a13, out13 := ALU1bit(typeConv(a[2]), typeConv(b[2]), out12, op)
	a14, out14 := ALU1bit(typeConv(a[1]), typeConv(b[1]), out13, op)
	a15, overflow := ALU1bitOF(typeConv(a[0]), typeConv(b[0]), out14, op)
	results := []uint8{typeConv2(a15), typeConv2(a14), typeConv2(a13), typeConv2(a12), typeConv2(a11), typeConv2(a10), typeConv2(a9), typeConv2(a8), typeConv2(a7), typeConv2(a6), typeConv2(a5), typeConv2(a4), typeConv2(a3), typeConv2(a2), typeConv2(a1), typeConv2(a0)}
	return results, overflow
}

//ALU16bitMult simulates multiplication
//----( 1 )--------
// if(MQ = 1){
// 	AC <- AC + MD
// } else {
// 	AC <- AC + 0
// }
//----( 2 )--------
// AC/MQ >> 1
func ALU16bitMult(a, b, c []byte) int {
	var MD, AC, MQ = a, c, b
	var zeroSet = c

	if typeConv(MQ[15]) == 1 {
		AC, _ = ALU16bit(AC, MD, 10)
	} else {
		AC, _ = ALU16bit(AC, zeroSet, 10)
	}

	var acTemp = typeConv2(typeConv(AC[15]))
	AC = append(AC[:0], AC[:15]...)
	var in = []uint8{0}
	AC = append(in, AC...)

	MQ = append(MQ[:0], MQ[:15]...)
	var to = []uint8{acTemp}
	MQ = append(to, MQ...)

	fmt.Println("MD: ", MD, " AC: ", AC, " MQ: ", MQ)
	return 0
}
