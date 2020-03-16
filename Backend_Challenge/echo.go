package Backend_Challenge

import (
  "encoding/csv"
  "fmt"
  "net/http"
  "strings"
)

func Echo(w http.ResponseWriter, r *http.Request) {
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
  for _, row := range records {
    fmt.Println(row)
    response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
  }
  fmt.Fprint(w, response)
}
