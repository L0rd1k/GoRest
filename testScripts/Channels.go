package main

import (
	"fmt"
	"time"
)

func sum(numArr []int, chnl chan func() (int, int)) {
	sum := 0
	mul := 1
	for _, elem := range numArr {
		sum += elem
		mul *= elem
	}
	chnl <- (func() (int, int) { return sum, mul })
}

func channelsBuffering(chnl *chan string) {
	*chnl <- "Hello"
	*chnl <- "World!"
}

func testFunction(chnl chan int) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	chnl <- 100
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	// // 1 Сhannels
	// messages := make(chan string)
	// go func() { messages <- "Run test message" }()
	// getMsg := <-messages
	// fmt.Println(getMsg)
	// ========================================================

	// // 2 Channels - return multiple values

	// testArr := []int{1, 2, 3, 4, 5}
	// chnl := make(chan func() (int, int))
	// go sum(testArr, chnl)
	// res_sum, res_mul := (<-chnl)()
	// fmt.Println("Sum: ", res_sum, " / Mul: ", res_mul)
	// ========================================================

	// // 3 Channels - Buffering

	// test_chnl := make(chan string, 2)
	// go channelsBuffering(&test_chnl)
	// fmt.Println(<-test_chnl, <-test_chnl)
	// ========================================================

	// // 4 Channels - Synchronization

	// new_test_chnl := make(chan int)
	// go testFunction(new_test_chnl)
	// <-new_test_chnl
	// ========================================================

	// // 5 Channel Directions

	// pings := make(chan string, 1)
	// pongs := make(chan string, 1)
	// ping(pings, "passed message")
	// pong(pings, pongs)
	// fmt.Println(<-pongs)
	// ========================================================

	// 6 Multiple channel opeartions - SELECT

	// chnl1 := make(chan string)
	// chnl2 := make(chan string)
	// go func() {
	// 	time.Sleep(7 * time.Second)
	// 	chnl1 <- "one"
	// }()
	// go func() {
	// 	time.Sleep(5 * time.Second)
	// 	chnl2 <- "two"
	// }()
	// for i := 0; i < 2; i++ {
	// 	select {
	// 	case msg1 := <-chnl1:
	// 		fmt.Println("received", msg1)
	// 	case msg2 := <-chnl2:
	// 		fmt.Println("received", msg2)
	// 	}
	// }
	// ========================================================

	// // 7 None-Blocking channel operations
	// messages := make(chan string)
	// signals := make(chan bool)

	// go func() {
	// 	messages <- "Hello"
	// }()

	// time.Sleep(time.Second)

	// select {
	// case msg := <-messages:
	// 	fmt.Println("received message", msg)
	// default:
	// 	fmt.Println("no message received")
	// }

	// msg := "hi"
	// select {
	// case messages <- msg:
	// 	fmt.Println("sent message", msg)
	// default:
	// 	fmt.Println("no message sent")
	// }

	// select {
	// case msg := <-messages:
	// 	fmt.Println("received message", msg)
	// case sig := <-signals:
	// 	fmt.Println("received signal", sig)
	// default:
	// 	fmt.Println("no activity")
	// }
	// ========================================================

	// 8 Close Channels

	channel_1 := make(chan int, 5)
	channel_2 := make(chan bool)

	go func() {
		for {
			j, more := <-channel_1
			if more {
				fmt.Println("Received channel_1", j)
			} else {
				fmt.Println("Received all channel_1")
				channel_2 <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		channel_1 <- j
		fmt.Println("sent channel_1", j)
	}

	close(channel_1)
	fmt.Println("sent all channel_1")
	<-channel_2
	// close(channel_1)

	// time.Sleep(time.Second)
}
