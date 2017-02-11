// Challenge 5) Write a program that creates N items like, where N is specified by a command line argument, where the key has a random value of 6 characters
// [{ value: "abcdef",  hex: "", base32: "" }, ...]
// enrich the items with hex and base32 based representation by processing each one in a go-routine

package main

import (
  "fmt"
  "math/rand"
  b32 "encoding/base32"
  hex "encoding/hex"
)

const letters = "abcdefghijklmnopqrstuvwxyz"

type UserInputMap struct {
  Key string `json:"key"`
  Base32 string `json:"base32"`
  Hex string `json:"hex"`
}

type Empty struct {}

func GenerateRandomString() (s string) {
  for len(s) < 6 {
    s += string(letters[rand.Intn(len(letters))])
  }
  return
}

func EncodeToB64(o *UserInputMap, c chan Empty) {
  bEnc := b32.StdEncoding.EncodeToString([]byte(o.Key))
  o.Base32 = bEnc
  c<- Empty{}
}

func EncodeToHex(o *UserInputMap, c chan Empty) {
  hEnc := hex.EncodeToString([]byte(o.Key))
  o.Hex = hEnc
  c<- Empty{}
}

func createUserInputMap() (arr []UserInputMap, err error ){
  var n int
  fmt.Print("Enter Number: ")
  _, err = fmt.Scanf("%d", &n)
  if err != nil {
    fmt.Println("Error: ", err)
    return nil, err
  }

  // create array with empty fields to be encoded
  arr = make([]UserInputMap, n)
  for i:=0; i<n; i++ {
    key := GenerateRandomString()
    arr[i] = UserInputMap{key, "", ""}
  }

  // encode keys
  c := make(chan Empty)
  for i, _ := range arr {
    go EncodeToB64(&arr[i], c)
    go EncodeToHex(&arr[i], c)
  }

  // clear channel
  for i:=0; i<n; i++ {
    _,_ = <-c, <-c
  }

  return arr, nil
}
