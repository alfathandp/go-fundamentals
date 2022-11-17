package main

import (
	"fmt"
	"strings"
)

const (
	str = "something"
	substr = "some"
)
func main()  {
	// 1. panjang kata dalam string
	sentence := "hello world"
	lenSentence := len(sentence)
	fmt.Println(lenSentence)

	// 2. compare string
	str1 := "abc"
	str2 := "abd"
	fmt.Println(str1 != str2)
	fmt.Println(str1 == str2)

	// 3. contain
	res := strings.Contains(str,substr)
	fmt.Println(res)

	// 4. mengambil character dari index tertentu
	value := "cat:dog"
	substring := value[4:7]
	fmt.Println(substring)

	// 5. replace
	Hewan := "katakwkwkwk"
	t := strings.Replace(Hewan,"a","o",-1)
	u := strings.Replace(Hewan, "a","o", 1)
	v := strings.Replace(Hewan,"k","p",3)
	fmt.Println(t,u,v)

	// 6. insert
	p:= "green"
	index := 2
	q := p[:index] + "hello" + p[index:]
	fmt.Println(p,q)

}