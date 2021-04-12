package main

import (
	"fmt"
	"time"
)

func testFunction(strValue string) {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(strValue, "...", i)
	}
}

func main() {
	fmt.Println("Running...")
	go testFunction("Asynchronouse call_1...")
	go testFunction("Asynchronouse call_2...")
	testFunction("Synchronouse call...")

	time.Sleep(time.Second)
}
