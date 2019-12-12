package main

import "fmt"

//通道关闭demo
func main() {
	ch1 := make(chan int, 2)
	//发送方
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("sender: sending element %d...\n", i)
			ch1 <- i
		}
		fmt.Println("sender: close the channel...")
		close(ch1)
	}()
	//接收方
	for {
		elem, ok := <-ch1
		if !ok {
			fmt.Println("receiver: closed channel")
			break
		}
		fmt.Printf("receiver: receiver an element: %d\n", elem)
	}
}
