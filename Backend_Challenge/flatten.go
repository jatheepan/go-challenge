package Backend_Challenge

import (
  "encoding/csv"
  "fmt"
  "net/http"
  "strings"
)

func Flatten(w http.ResponseWriter, r *http.Request) {
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
  for index, row := range records {
    separator := ","
    if (index +1) % len(records) == 0 {
      separator = ""
    }
    response = fmt.Sprintf("%s%s%s", response, strings.Join(row, ","), separator)
  }
  fmt.Fprint(w, response)
}
