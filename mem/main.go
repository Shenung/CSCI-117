package main

import "fmt"

//lw  100011
//sw  101011

var reg = [8]int{}
var mainMem = [128]int{} //init address plus 5

//this function starts before main programs is ran to initialize mainMem values
func init() {
	for i := range mainMem {
		mainMem[i] = i + 5
	}
}

//make struct of row and block and put into array
type block struct { //[valid][tag][data]
	valid bool
	tag   int
	data  string
}

type row struct {
	block1 [8]block
	block2 [8]block
}

//this is a cache that is type row used for storing the data
//simulates a cache memory block
var cache row

func main() {
	var code = make([]byte, 32)
	fmt.Println("32bit code -> ")
	fmt.Scanln(&code)
	var convertedCode = string(code)
	decode(convertedCode)
}

func decode(code string) {
	var op = code[:6]
	var rs = code[6:11]
	var rt = code[11:16]
	var offset = code[16:32]
	fmt.Println("op code-> ", op)
	fmt.Println("rs field-> ", rs)
	fmt.Println("rt field-> ", rt)
	fmt.Println("offset field-> ", offset)
	instruction(op, rs, rt, offset)
}

func instruction(op, rs, rt, offset string) {
	if op == "100011" {

	} else if op == "101011" {

	}
}
