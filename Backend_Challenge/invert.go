package Backend_Challenge

import (
  "encoding/csv"
  "fmt"
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

  var response string
  colIndex, rowIndex := 0, 0
  for colIndex < 3 {
    response += records[rowIndex][colIndex]
    if rowIndex < 2 {
      response += ","
    }
    rowIndex += 1
    if rowIndex >= 3 {
      response += "\n"
      colIndex += 1
      rowIndex = 0
    }
  }

  fmt.Fprint(w, response)
}
