// Challenge 6) Write a program that writes the output of exercise 5 to a json file

package main

// To run in terminal from root directory:
// go run UserInputMap/UserInputMap.go UserInputMap/Main.go

import (
  "fmt"
  "io/ioutil"
  "encoding/json"
)

func main() {
  usermap, err := createUserInputMap()
  if err != nil {
    fmt.Println("Error Creating UserInputMap: ", err)
    return
  }
  usermapjson, _  := json.Marshal(usermap)
  err = ioutil.WriteFile("UserInputMap/output.json", usermapjson, 0644)
  if err != nil {
    fmt.Println("Error Writing File: ", err)
    return
  }
  fmt.Println("New file output.json created")
}
