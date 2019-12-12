package demo1

import (
	"fmt"
	"math/rand"
)

func main() {
	//只能发不能收的通道
	uselessChan := make(chan<- int, 1)
	//只能收不能发的通道
	uselessChan2 := make(<-chan int, 1)
	//打印两个通道的指针的16进制表示
	fmt.Printf("the useless channels: %v, %v\n", uselessChan, uselessChan2)
}

//示例2
func SendInt(ch chan<- int) {
	ch <- rand.Intn(1000)
}

//示例3
type Notifier interface {
	SendInt(ch chan<- int)
}

//示例4
func getIntChan() <-chan int {
	num := 5
	ch := make(chan int, num)
	for i := 0; i < num; i++ {
		ch <- i
	}
	close(ch)
	return ch
}

//示例5
type GetIntChan func() <-chan int
