package main

import "fmt"

func addUp(x int, y int) int {
  return x + y
}

func anotherFunction(f func() string) string {
  return f()
}

func main () {
  //basic function
  fmt.Printf("%v\n", addUp(1,5))

  //passing function as values
  fn := func() {
    fmt.Println("function called")
  }

  fn()

  // passing function as argument
  afn := func() string {
    return "function call"
  }

  fmt.Println(anotherFunction(afn))

}


//engine light the purge valve sillinoed

//203 8821971 mel gomez, part of the emition
