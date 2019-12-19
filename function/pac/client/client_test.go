package client

import (
	"function/pac/server"
	"testing"
)

func main() {

}

//小写的方法声明在包外不能被访问
func TestPac(t *testing.T) {
	t.Log(_server.GetFibonacci(10))
}

func TestAAA(t *testing.T) {

}
