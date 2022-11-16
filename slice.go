package main

import "fmt"

func main() {

	// append & copy
	var colors = []string{"red","blue","green"}
	colors = append (colors,"yellow")
	copied_colors := make([]string,100)

	copy(copied_colors, colors)
	fmt.Println(copied_colors)
}
