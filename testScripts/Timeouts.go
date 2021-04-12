package main

import (
	"fmt"
	"time"
)

func main() {
	test_channel1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		test_channel1 <- "result 1"
	}()

	select {
	case res := <-test_channel1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

}
