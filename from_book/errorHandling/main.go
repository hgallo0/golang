package main

import (
  "fmt"
  "io/ioutil"
  "errors"
)

func main() {
  err := errors.New("Something went wrong")
  file, err := ioutil.ReadFile("foo.txt")

   if err != nil {
     fmt.Println(err)
     return
   }

   //fmt.Println(file)
   fmt.Println("%s", file)
}
