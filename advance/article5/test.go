package main

import "fmt"

var block = "package"

func main() {
	var block = "package"
	{
		block := "inner"
		fmt.Println(block)
	}
	fmt.Println(block)
}
