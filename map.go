package main

import "fmt"

func main()  {
	//deklarasi map
	harga := map[string]int{"bakso": 15000, "somay": 1000}
	warna := map[string]string{"tembok": "kuning", "lantai":"putih"}
	fmt.Println(harga, warna)

	var gorengan = make(map[string]int)
	gorengan["bakwan"] = 1000
	fmt.Println(gorengan)
	
}