package Backend_Challenge

import (
  "encoding/csv"
  "fmt"
  "math"
  "net/http"
)

func Invert(w http.ResponseWriter, r *http.Request) {
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

  recordsCount := math.Pow(float64(len(records)), 2)
  var response string
  for i := float64(0); i < recordsCount; {
    separator := ","
    colIndex := int(math.Floor(i / 3))
    rowIndex := int(i) % len(records)
    if rowIndex == len(records) -1 {
      separator = "\n"
    }
    item := records[rowIndex][colIndex]
    response = fmt.Sprintf("%s%s%s", response, item, separator)
    i += 1
  }

  fmt.Fprint(w, response)
}
