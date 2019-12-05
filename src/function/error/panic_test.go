package error

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

func TestExit(t *testing.T) {
	fmt.Println("hello world")
	os.Exit(-1)
}

//panic会输出调用栈的信息
func TestPanic(t *testing.T) {
	defer func() {
		fmt.Println("end---")
	}()
	fmt.Println("hello world")
	panic(errors.New("error---"))
}

//恢复错误
func TestRecover(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover---")
		}
	}()
	fmt.Println("hello world")
	panic(errors.New("error---"))
}
