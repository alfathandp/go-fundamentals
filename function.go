package main

import (
	"fmt"
	"math"
)

func main()  {
	hour := 14
	greeting(hour)
	sayHello()

	var side = 5
	wide := calculateSquare(side)
	fmt.Println("luas persegi : ", wide)

	var diameter float64 = 15
	keliling, luas := calculateCircle(diameter)
	fmt.Printf("luas lingkaran : %.3f \n", luas)
	fmt.Printf("keliling lingkaran: %.2f \n", keliling)

	m := multiplication(5,5)
	fmt.Println("5x5 = ", m)
}

//deklarasi fungsi tanpa parameter
func sayHello()  {
	fmt.Println("hellllloooo")
}

// dekkarasi fungsi dengan parameter
func greeting(hour int) {
	if hour < 12 {
		fmt.Println("selamat pagi dunia")
	} else if hour < 18 {
		fmt.Println("seamat sore dunya")
	} else {
		fmt.Println("selamat malam")
	}
}

// deklarasi fungsi dengan single return value
func calculateSquare(side int) int{
	return side*side
}

// deklarasi multiple return value
func calculateCircle(diameter float64) (float64,float64)  {
	var keliling = math.Pi * math.Pow(diameter/2,2)
	var luas = math.Pi * diameter

	// return 2 value
	return keliling, luas
}

// fungsi yang memiliki nama return value
func multiplication(a,b int) (mul int) {
	mul = a*b
	return
}