package main

func deadlockChannel() {
	ch1 := make(chan int)
	//fatal error: all goroutines are asleep - deadlock!
	ch1 <- 13
	n := <-ch1
	println(n)
}

func normalChannel() {
	ch1 := make(chan int)
	go func() {
		//将发送操作放入一个新goroutine中执行
		ch1 <- 13
	}()
	n := <-ch1
	println(n)
}
func main() {
	//deadlockChannel()
	normalChannel()
}
