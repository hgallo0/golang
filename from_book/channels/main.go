package main

import (
	"fmt"
	"time"
)

func slowFunc(c chan string, someValue string) {
	/*for msg := range c {
	  fmt.Println(msg)
	}*/
	time.Sleep(time.Second * 2)
	c <- someValue
}

func ping1(c chan string) {
	time.Sleep(time.Second * 3)
	c <- "ping on channel1"
}

func ping2(c chan string) {
	time.Sleep(time.Second * 2)
	c <- "ping on channel2"
}

func main() {
	messages := make(chan string)
	go slowFunc(messages, "Hello")

	msg := <-messages
	fmt.Println(msg)

	//select
	channel1 := make(chan string)
	channel2 := make(chan string)

	go ping1(channel1)
	go ping2(channel2)

	select {
	case msg1 := <-channel1:
		fmt.Println("received", msg1)
	case msg2 := <-channel2:
		fmt.Println("received", msg2)
	}
}
