package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(10 * time.Second)

	// <-timer1.C
	fmt.Println("Timer1 1 fired")

	timer2 := time.NewTimer(5 * time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	<-timer1.C
	// stop2 := timer2.Stop()
	// if stop2 {
	// 	fmt.Println("Timer 2 stopped")
	// }
	time.Sleep(2 * time.Second)
}
