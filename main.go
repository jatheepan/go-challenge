package main

import (
  "league.com/challenge/Backend_Challenge"
  "log"
  "net/http"
)

func main() {
  http.HandleFunc("/echo", Backend_Challenge.Echo)
  http.HandleFunc("/invert", Backend_Challenge.Invert)
  http.HandleFunc("/sum", Backend_Challenge.Sum)
  http.HandleFunc("/multiply", Backend_Challenge.Multiply)
  log.Fatal(http.ListenAndServe(":8080", nil))
}
