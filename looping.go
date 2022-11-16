package main

import "fmt"

func main()  {
	// looping for
	for i := 0; i < 5; i++ {
		fmt.Println("#")
	}

	// looping with continue & break
	for i := 0; i < 5; i++ {
		if i == 1 {
			continue // memaksa lanjut ke perulangan selanjutnya
		}
		if i > 3 {
			break // memaksa perulangan berhenti
		}
		fmt.Println(i)
	}

	//looping over string
	var sentences string = "hello"
	for i := 0; i < len(sentences); i++ {
		fmt.Printf(string(sentences[i]) + "-")
	}
	fmt.Println("=========")

	for a,b := range sentences {
		fmt.Printf("huruf %c ada di posisi ke-%d \n", b, a)
	}
}