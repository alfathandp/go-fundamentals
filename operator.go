package main

import "fmt"

func main() {
	// + adalah operator, 1 dan 2 adalah operand 
	x:= 1+2
	fmt.Println(x)

	//expression
	a := 5
	b := 6
	c := a*b
	fmt.Println(c)

	//luas segitiga
	alas := 5
	tinggi := 10
	luasSegitiga := (alas*tinggi)/2
	fmt.Println(luasSegitiga)

	//operasi string
	helloWorld := "hello" + " " + "dunia tipu tipu"
	fmt.Println(helloWorld)

}