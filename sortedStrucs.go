// Challenge 4) Create an array of random key-value items (structs), like:
// [{ key: "c", value: "123" }, { key: "a", value: "456" }, { key: "b", value: "789" }]
// and sort the by key and by value

package main

import (
  "fmt"
  "math/rand"
  "strconv"
  "sort"
)

const arrayLength = 10
const maxIntValue = 100
const letters = "abcdefghijklmnopqrstuvwxyz"

type MyObject struct {
  Key string
  Value string
}
type ByKeys []MyObject
type ByValues []MyObject

// sort functions
// sorts numbers and letters lexicographically

func (m ByKeys) Len() int {
    return len(m)
}
func (m ByKeys) Swap(i, j int) {
    m[i], m[j] = m[j], m[i]
}
func (m ByKeys) Less(i, j int) bool {
    return m[i].Key < m[j].Key
}

func (m ByValues) Len() int {
    return len(m)
}
func (m ByValues) Swap(i, j int) {
    m[i], m[j] = m[j], m[i]
}
func (m ByValues) Less(i, j int) bool {
    return m[i].Value < m[j].Value
}

// random value generators
func GenerateRandomLetter() string {
  return string(letters[rand.Intn(len(letters))])
}

func GenerateRandomNumber() string {
  return strconv.Itoa(rand.Intn(maxIntValue))
}

//main
func main() {
  arr := make([]MyObject, arrayLength)

  for i:=0; i<arrayLength; i++ {
    randLetter := GenerateRandomLetter()
    randNumber := GenerateRandomNumber()
    arr[i] = MyObject{randLetter, randNumber}
  }

  sort.Sort(ByKeys(arr))
  fmt.Println("Sorted by keys: ", arr)

  sort.Sort(ByValues(arr))
  fmt.Println("\nSorted by values: ", arr)
}
