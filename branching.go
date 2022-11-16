package main

import "fmt"

func main() {
	// if , if else , else
	nama := "fathan"
	hour := 2
	if hour < 12 {
		fmt.Println("selamat pagi" + " " + nama )
	} else if hour > 12 && hour < 18 {
		fmt.Println("selamat siang")
	} else {
		fmt.Println("selamat malam")
	}

	// switch
	today := 3
	switch today {
	case 1:
		fmt.Println("today is monday")
	case 2:
		fmt.Println("today is selasa")
	case 3:
		fmt.Println("today is wednesday")
	default:
		fmt.Println("invalid day")
	}
}