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

func main() {
    messages := make(chan string)
    go slowFunc(messages, "Hello")

    msg := <-messages
    fmt.Println(msg)
}
