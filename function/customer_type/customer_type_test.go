package customer_type

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type IntConv func(op int) int

//方法时长计算
func timeSpent(inner func(op int) int) IntConv {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Print("timeSpent: ", time.Since(start).Seconds())
		return ret
	}
}

func returnMultiVal() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func TestFn(t *testing.T) {
	a, _ := returnMultiVal()
	t.Log(a)
	spent := timeSpent(timeSlow)
	t.Log(spent(10))
	t.Log(sum(1, 2, 3, 4, 5))
}

func timeSlow(op int) int {
	time.Sleep(time.Second)
	return op
}

//可变参数测试
func sum(ops ...int) int {
	s := 0
	for _, op := range ops {
		s += op
	}
	return s
}
