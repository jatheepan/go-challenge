package main

import (
  "encoding/csv"
  "errors"
  "fmt"
  "league.com/challenge/Backend_Challenge"
  "log"
  "net/http"
  "strings"
)

func ParseFile(r *http.Request) (records [][]string, err error) {
  file, fileHeader, err := r.FormFile("file")
  // Validate file
  if file == nil {
    return nil, errors.New("A CSV File is required")
  }
  if err != nil {
    return nil, err
  }
  defer file.Close()
  fileNames := strings.Split(fileHeader.Filename, ".")
  fileExtension := strings.ToLower(fileNames[len(fileNames) -1])
  // Validate file format
  if fileExtension != "csv" {
    return nil, errors.New("A CSV File is required")
  }

  return csv.NewReader(file).ReadAll()
}

func Handle(manipulatorFn func([][]string) string) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    // Allow only POST request as we are accepting a file
    if r.Method != http.MethodPost {
      w.WriteHeader(http.StatusMethodNotAllowed)
      w.Write([]byte("Method not Allowed"))
      return
    }
    records, err := ParseFile(r)
    if err != nil {
      w.WriteHeader(http.StatusUnprocessableEntity)
      w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
      return
    }
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

