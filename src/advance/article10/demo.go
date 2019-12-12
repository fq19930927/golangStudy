package main

//演示几种通道发送方和接收方都会被阻塞的demo
func main() {
	//示例1
	ch1 := make(chan int, 1)
	ch1 <- 1
	//ch1 <- 2 测试通道已满，会造成阻塞

	//示例2
	ch2 := make(chan int, 2)
	//elem, ok:= <-ch2 通道已空,会造成接收阻塞
	ch2 <- 1

	//示例3
	var ch3 chan int
	//ch3 <- 1 通道为nil,会永久阻塞
	// <- ch3 通道为nil,会永久阻塞
	_ = ch3
}
