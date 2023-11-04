package main

import (
	"fmt"
	"strconv"
	"time"
)

func makeCakeAndSend(cs chan string, count int) {
	for i := 1; i <= count; i++ {
		time.Sleep(1 * time.Second)
		cakeName := "Strawberry Cake " + strconv.Itoa(i)
		fmt.Println(">>>>>>>>>>>>> send cake: ", cakeName)
		cs <- cakeName //send a strawberry cake
	}
	fmt.Println("close chan...")
	//因为消费者比生产者处理慢，在消费者没消费完之前就close了。
	//从效果上看，调用close 之后，chan中的消息还是能够顺利被消费完的
	close(cs)
}

func receiveCakeAndPack(cs chan string) {
	for s := range cs {
		//假设消费者处理比生产者慢
		time.Sleep(2 * time.Second)
		fmt.Println("<<<<<<<<<<<<<<< received cake: ", s)
	}
}

func main() {
	cs := make(chan string)
	go makeCakeAndSend(cs, 5)
	go receiveCakeAndPack(cs)

	time.Sleep(20 * time.Second)
}
