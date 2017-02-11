// Challenge 3) Write a simple HTTP server, return JSON object with current time, like
// {
// 	"time": "2006-01-02T15:04:05Z07:00"
// }
package main

import (
  "fmt"
  "net/http"
  "time"
  "encoding/json"
)

type TimeResponse struct {
  Time time.Time `json:"time"`
}

func handler(w http.ResponseWriter, r *http.Request) {
  response := &TimeResponse{ time.Now()}
  slc, _ := json.Marshal(response)
  fmt.Fprintf(w, "%s", slc)
}

func main() {
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8080", nil)
}
