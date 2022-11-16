package main

import "fmt"

func main() {
	// deklarasi array short and long
	x := [5] int {1,2,3,5,5}
	var y [5]int = [5]int{1,2,3}
	z :=[3]string {"saya","adalah"}

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	primes := [] int {2,3,5}
	//looping array cara pertama
	for index :=0; index <len(primes); index++ {
		fmt.Println(primes[index])
	}

	// looping array cara kedua
	for index, elements := range primes {
		fmt.Println(index, "==>", elements)
	}

	// looping array cara ketiga
	index :=0
	for range primes {
		fmt.Println(primes[index])
		index++
	}
}