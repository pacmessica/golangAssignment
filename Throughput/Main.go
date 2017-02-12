// Challenge 7) Write a program with 3 go-routines, which does the following:
// - Go routine 1 generates continuous stream of items, which are sent to GR2. The time between sending items varies between 10ms and 100ms
// - Go routine 2 makes micro batches, either when it receives 25 items or 2 seconds have passed
// - Go routine 3 reports the throughput to screen

package main

import (
  "fmt"
  "time"
  "math/rand"
  "strings"
)

type Item struct {
  Value int
}

// go routine 1
func generateItems(out chan Item) {
  for {
      x := (rand.Intn(100))
      out <- Item{x}
      t := (rand.Intn(90) + 10) // create random int between 10-100
      time.Sleep(time.Duration(t) * time.Millisecond)
  }
}

// go routine 2
func batchItems(in chan Item, out chan []Item) {
  start := time.Now()
  arr := make([]Item, 0)

  for n := range in {
    arr = append(arr, n)

    for time.Since(start) > time.Second * 2 || len(arr) == 25 {
      out <- arr
      arr = arr[len(arr):]
      start = time.Now()
    }
  }
}

// go routine 3
func reportThroughput(in chan []Item, out chan []Item) {
  tick := time.Tick(time.Second)
  backspaces := ""
  throughput := 0
  for {
    select {
    case <-tick:
        throughputStr := fmt.Sprintf("%d batches/s", throughput)
        fmt.Printf(backspaces + throughputStr)
        throughput = 0
        backspaces = strings.Repeat("\b", len(throughputStr))
    case val := <-in:
      out <- val
      throughput += 1
    default:
    }
  }
}

func main() {
    chItems := make(chan Item, 5)
    chBatchToBeObserved := make(chan []Item, 5)
    chBatch := make(chan []Item, 5)
    go generateItems(chItems)
    go batchItems(chItems, chBatchToBeObserved)
    go reportThroughput(chBatchToBeObserved, chBatch)


    // consume channel
    for n := range chBatch {
      _ = n
    }
}
