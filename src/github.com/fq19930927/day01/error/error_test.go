package error

import (
	"errors"
	"fmt"
	"testing"
)

var LessThanZero = errors.New("num < 0 error")
var MoreThanHundred = errors.New("num > 100 error")

//异常编写
func GetFibonacci(n int) ([]int, error) {
	if n < 0 {
		return nil, LessThanZero
	} else if n > 100 {
		return nil, MoreThanHundred
	}
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func TestFib(t *testing.T) {
	if v, err := GetFibonacci(-10); err != nil {
		if err == LessThanZero {
			fmt.Println("hehehe")
		} else {
			t.Error(err)
		}
	} else {
		t.Log(v)
	}
}
