package main

import (
  "encoding/csv"
  "errors"
  "fmt"
  "league.com/challenge/Backend_Challenge"
  "log"
  "net/http"
)

func Handle(manipulatorFn func([][]string) string) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
      w.WriteHeader(http.StatusMethodNotAllowed)
      return
    }
    file, _, err := r.FormFile("file")
    if file == nil {
      err = errors.New("File is required")
    }
    if err != nil {
      w.WriteHeader(http.StatusUnprocessableEntity)
      w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
      return
    }
    defer file.Close()
    records, err := csv.NewReader(file).ReadAll()
    w.Write([]byte(manipulatorFn(records)))
  })
}

func main() {
  http.Handle("/echo", Handle(Backend_Challenge.Echo))
  http.Handle("/invert", Handle(Backend_Challenge.Invert))
  http.Handle("/flatten", Handle(Backend_Challenge.Flatten))
  http.Handle("/sum", Handle(Backend_Challenge.Sum))
  http.Handle("/multiply", Handle(Backend_Challenge.Multiply))
  log.Fatal(http.ListenAndServe(":8080", nil))
}

