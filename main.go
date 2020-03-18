package main

import (
  "encoding/csv"
  "errors"
  "fmt"
  "league.com/challenge/Backend_Challenge"
  "log"
  "net/http"
)

func Validate(next http.Handler) http.Handler {
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
    next.ServeHTTP(w, r)
  })
}

func main() {
  handler := func(fn func([][]string) string) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
      file, _, err := r.FormFile("file")
      defer file.Close()
      if err != nil {
        w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
        return
      }
      records, err := csv.NewReader(file).ReadAll()
      w.Write([]byte(fn(records)))
    })
  }

  http.Handle("/echo", Validate(handler(Backend_Challenge.Echo)))
  http.Handle("/invert", Validate(handler(Backend_Challenge.Invert)))
  http.Handle("/flatten", Validate(handler(Backend_Challenge.Flatten)))
  http.Handle("/sum", Validate(handler(Backend_Challenge.Sum)))
  http.Handle("/multiply", Validate(handler(Backend_Challenge.Multiply)))
  log.Fatal(http.ListenAndServe(":8080", nil))
}

