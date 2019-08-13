package main

import (
  "fmt"
)

func main() {
  //arrays
  var cheeses [2]string

  cheeses[0] = "American"
  cheeses[1] = "Mozzarella"

  fmt.Println(cheeses)

  //slices
  var cheeses_s = make([]string, 2)
  cheeses_s[0] = "American"
  cheeses_s[1] = "Mozzarella"

  cheeses_s = append(cheeses_s, "Camembert")

  fmt.Println(cheeses_s)

  //maps
  var players = make(map[string]int)

  players["cook"] = 32

  players["bairstow"] = 27

  players["stokes"] = 26

  fmt.Println(players)
  fmt.Println(players["cook"])

  delete(players, "cook")
  fmt.Println(players)
}
