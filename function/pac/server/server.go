package _server

import (
	"fmt"
	"testing"
)

func init() {
	fmt.Println("init1 ----")
}

//被外部调用时,所有init方法都会执行
func init() {
	fmt.Println("init2 ----")
}

func GetFibonacci(n int) ([]int, error) {
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func TestAAA(t *testing.T) {

}
