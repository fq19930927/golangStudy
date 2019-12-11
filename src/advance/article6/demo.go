package main

import "fmt"

var container = []string{"tony", "lili"}

func main() {
	container := map[int]string{1: "asd", 2: "ddd"}
	if _, ok := interface{}(container).(map[int]string); ok {
		fmt.Println("type is right")
	} else {
		fmt.Println("type is wrong")
	}
}
