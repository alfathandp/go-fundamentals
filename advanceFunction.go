package main

import "fmt"

// variadic function
func sum(numbers ...int) (int, int) {
	var total int = 0
	var indexx int = 0
	for index, number:= range numbers {
		total+= number
		indexx += index
	}
	return total , indexx
}


func main() {

	defer func ()  {
		fmt.Println("boong deng yg ini yg terakhir")
	}()

	defer func ()  {
		fmt.Println("ini fungsi defer, akan dijalankan terakhir")
	}()
	
	avg,index := sum (2,3,4,5)
	fmt.Println(avg,index)

	//anonymous function 
	func ()  {
		fmt.Println("selamat datang")
	}()

	// dimasukkan kedalam variable
	greet := func ()  {
		fmt.Println("ini adalah contoh anonymous function")
	}
	greet()

	// anonymous function yang memiliki parameter
	func (sentence string)  {
		fmt.Println(sentence)
	}("halo dunia")
}
