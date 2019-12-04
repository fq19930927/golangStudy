package _interface_test

import (
	"fmt"
	"testing"
)

func DoSomething(p interface{})  {
	/*if i,ok:=p.(int);ok{
		fmt.Println("integer ", i)
		return
	}
	if s, ok:=p.(string);ok {
		fmt.Println("string", s)
		return
	}
	fmt.Println("unknow type")*/
	switch v:= p.(type) {
	case int:
		fmt.Println("integer ", v)
	case string:
		fmt.Println("string", v)
	default:
		fmt.Println("unknow type")
	}
}

func TestEmptyInterface(t *testing.T)  {
	DoSomething(10)
	DoSomething("10")
}
