package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	example1()
	example2()
}

func example1() {
	//准备通道
	intChannels := [3]chan int{
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
	}
	//随机选择一个通道并向它发送元素值
	index := rand.Intn(2)
	fmt.Printf("the index: %v\n", index)
	intChannels[index] <- index
	//哪一个通道有可取的元素值,哪个分支就会被执行
	select {
	case <-intChannels[0]:
		fmt.Println("the first is selected")
	case <-intChannels[1]:
		fmt.Println("the second is selected")
	case elem := <-intChannels[2]:
		fmt.Printf("the third is selected, elem is %d\n", elem)
	default:
		fmt.Println("none is selected")
	}
}

//示例2
func example2() {
	intChan := make(chan int, 1)
	//一秒后关闭通道
	time.AfterFunc(time.Second, func() {
		close(intChan)
	})
	select {
	case _, ok := <-intChan:
		if !ok {
			fmt.Println("the chan is closed")
			break
		}
		fmt.Println("the chan is selected")
	}
}
