package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main start================")
	done := make(chan bool)
	go dispatch(done)
	time.Sleep(1 * time.Second)

	done <- true
	fmt.Println("main done <- true---------")
	<-done //相当于异步转同步
	fmt.Println("main end done================")

}

func dispatch(done chan bool) {
	fmt.Println("dispatch start=====================")
	select {
	case <-done:
		defer func() {
			fmt.Println("dispatch defer done <- true-------------")
			done <- true
		}()
		fmt.Println("dispatch receive done--------------")
		time.Sleep(time.Second)
	}
	fmt.Println("dispatch end=====================")
}
