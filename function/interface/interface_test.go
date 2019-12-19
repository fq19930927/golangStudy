package _interface

import (
	"fmt"
	"testing"
)

type Programmer interface {
	WriteHelloword() string
}

type GoProgrammer struct {
}

func (g *GoProgrammer) WriteHelloword() string {
	return "go helloworld"
}

type JavaProgrammer struct {
}

func (j *JavaProgrammer) WriteHelloword() string {
	return "java helloworld"
}

func writeFirstProgrammer(programmer Programmer) {
	fmt.Printf("%T %v\n", programmer, programmer.WriteHelloword())
}

func TestDuotai(t *testing.T) {
	goP := new(GoProgrammer)
	javaP := new(JavaProgrammer)
	writeFirstProgrammer(goP)
	writeFirstProgrammer(javaP)
}

func TestInterface(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloword())
}
