package main

import (
  "fmt"
  "strconv"
)

type Movie struct {

    Name string

    Rating float64

}

type Address struct {
    Street string
    City string
    State string
    ZipCode int
}


func (m *Movie) summary() string {

  r := strconv.FormatFloat(m.Rating, 'f', 1, 64)

  return m.Name + ", " + r

}

func (a *Address) summary() string {


  z := strconv.FormatInt(int64(a.ZipCode), 8)
  //zip := ("%05d", z)

  return a.Street + ", " + a.City + ", " + a.State + ", " + z
//https://socketloop.com/tutorials/golang-convert-octal-value-to-string-to-deal-with-leading-zero-problem
}

func recursionCountDigits(number int) int {
 if number < 10 {
  return 1
 } else {
  return 1 + recursionCountDigits(number/10)
 }
}


func main() {

    m := Movie{

        Name:   "Spiderman",
        Rating: 3.2,

    }

    a := Address{
      Street: "somewhere",
      City: "LouisVile",
      State: "CA",
      ZipCode: 06563,
    }

    fmt.Println(m.summary())
    fmt.Println(a.summary())

}
