package Backend_Challenge

import (
  "encoding/csv"
  "fmt"
  "net/http"
  "strconv"
)

func Multiply(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodPost {
    w.WriteHeader(http.StatusMethodNotAllowed)
    return
  }
  file, _, err := r.FormFile("file")

  if err != nil {
    w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
    return
  }
  defer file.Close()
  records, err := csv.NewReader(file).ReadAll()
  if err != nil {
    w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
    return
  }
  var multiply float64 = 1
  for _, row := range records {
    for _, item := range row {
      f, err := strconv.ParseFloat(item, 64)
      if err == nil {
        multiply *= f
      }
    }
  }
  fmt.Fprint(w, multiply)
}
