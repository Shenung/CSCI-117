//Author: Shenung Fouamvung

package main

//MIPS processor simulation writen in Go Lang
//must have Go Lang installed to compile and run
//takes in 2 16bit binary values and computes them using logic gates to come to a result
//to run, CD to the file directory and use the command in command line terminal "go run main.go"
//example inputs:
//	a: 000000000001010
//  b: 000000000001101
//output:
//	result: 0000000000010111
//	overflow: 0

import (
	"CSCI-117/project/components"
	"fmt"
)

func main() {
	var a, b, ac = make([]byte, 16), make([]byte, 16), make([]byte, 16)

	fmt.Println("Enter 16bit binary [A]: ")
	fmt.Scanln(&a)
	fmt.Println("Enter 16bit binary [B]: ")
	fmt.Scanln(&b)
	fmt.Println("Enter 16bit AC value: ")
	fmt.Scanln(&ac)

	fmt.Println(a)
	fmt.Println(b)

	results, overflow := components.ALU16bit(a, b, 10)
	for i := range results {
		fmt.Print(results[i])
	}
	fmt.Println()
	fmt.Println("overflow:", overflow)

	fmt.Println(components.ALU16bitMult(a, b, ac))

}
