package main

import (
  "fmt"
  "os"
  "flag"
)

func main() {

  s := flag.String("s", "Hello world", "String help text")
  flag.Parse()
  fmt.Println("value of s:", *s)
  
  for i, arg := range os.Args {
    fmt.Println("argument", i, "is", arg)
  }


}
